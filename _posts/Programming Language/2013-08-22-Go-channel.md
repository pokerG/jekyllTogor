---
date: 2013-08-22
layout:  post
title:  channel的两个特性
tagline:  "不要通过共享内存来通信，而应该通过通信来共享内存。"
categories:
-  Programming Language
tags:
- Go
---

&emsp;&emsp;channel是Go提供的语言级goroutine间的通信方式。它提供了一种优雅而又强大的，在不使用锁的情况下，从某个goroutine向其他goroutine发送数据流的方式。

&esmp;emsp;这次着重讨论channel的两个重要特性。参考：[http://dave.cheney.net/2013/04/30/curious-channels](http://dave.cheney.net/2013/04/30/curious-channels)

被关闭的channel不会block
---
&esmp;emsp;一个channel一旦被关闭，就不能再向其发送数据，但是仍然可以获取其中的。

	package main
 
	import "fmt"
 
	func main() {
        ch := make(chan bool, 2)
        ch <- true
        ch <- true
        close(ch)
 
        for i := 0; i < cap(ch) +1 ; i++ {
                v, ok := <- ch
                fmt.Println(v, ok)
        }
	}

我们创造了缓冲区为两个值的channel。<br/>
输出结果:

	true true
	true true

这个程序首先输出发送到channel的两个值，然后在第三次<-时返回false和false。第一个false是channel类型的零值。第二个表示channel的启用状态。

能够检测channel 的启用状态也是很有用的一个特性。可用于对channel进行range操作。

&esmp;emsp;我们在这里举一个close channel的有用例子。首先看这个程序：

	package main
 
	import (
        "fmt"
        "sync"
        "time"
	)
 
	func main() {
        finish := make(chan bool)
        var done sync.WaitGroup
        done.Add(1)
        go func() {
                select {
                case <-time.After(1 * time.Hour):
                case <-finish:
                }
                done.Done()
        }()
        t0 := time.Now()
        finish <- true // 发送关闭信号
        done.Wait()    // 等待 goroutine 结束
        fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
	}

这个程序能够正常结束 `Waited 129.607us for goroutine to stop`

&emsp;&emsp;但是存在一些问题。首先finish channel 是不带缓冲的。如果接收方忘记了在select语句中添加finish，向其发送数据可能会导致阻塞。当然可以通过对要发送到的select块进行封装，以确保不会阻塞，或者设置带有缓冲的channel。但是，如果有许多goroutine都监听在finish channel上，那就需要跟踪这个情况，并且记住发送正确数量的数据给finish channel。如果无法控制goroutine的创建会很棘手；同时它们也可能是由程序的另一部分创建的，例如在响应网络请求的时候。

&emsp;&emsp;对于这种情况，我们可以利用已经关闭的channel会实时返回这一机制。

	package main
 
	import (
        "fmt"
        "sync"
        "time"
	)
 
	func main() {
        const n = 100
        finish := make(chan bool)
        var done sync.WaitGroup
        for i := 0; i < n; i++ { 
                done.Add(1)
                go func() {
                        select {
                        case <-time.After(1 * time.Hour):
                        case <-finish:
                        }
                        done.Done()
                }()
        }
        t0 := time.Now()
        close(finish)    // 关闭 finish 使其立即返回
        done.Wait()      // 等待所有的 goroutine 结束
        fmt.Printf("Waited %v for %d goroutines to stop\n", time.Since(t0), n)
	}

&emsp;

	Waited 231.385us for 100 goroutines to stop

&emsp;&emsp;那么这里发生了什么？当 finish channel 被关闭后，它会立刻返回。那么所有等待接收 time.After channel 或 finish 的 goroutine 的 select 语句就立刻完成了，并且 goroutine 在调用 done.Done() 来减少 WaitGroup 计数器后退出。这个强大的机制在无需知道未知数量的 goroutine 的任何细节而向它们发送信号而成为可能，同时也不用担心死锁。

nil channel 永远都是 block
----

&emsp;&emsp;再看Unknwon的《Go fundamental programming》时，他提到过一个类似于这样的例子。

	// WaitMany 等待 a 和 b 关闭。
	func WaitMany(a, b chan bool) {
        var aclosed, bclosed bool
        for !aclosed || !bclosed {
                select {
                case <-a:
                        aclosed = true
                case <-b:
                        bclosed = true
                }
        }
	}

&emsp;&emsp;WaitMany() 用于等待 channel a 和 b 关闭是个不错的方法，但是有一个问题。假设 channel a 首先被关闭，然后它会立刻返回。但是由于 bclosed 仍然是 false，程序会进入死循环，而让 channel b 永远不会被判定为关闭。

&emsp;&emsp;当时他也没用给出什么好的解决办法。在这里解决这个问题的比较好的办法就是利用nil channel的阻塞特性。

	package main
 
	import (
        "fmt"
        "time"
	)
 
	func WaitMany(a, b chan bool) {
        for a != nil || b != nil {
                select {
                case <-a:
                        a = nil 
                case <-b:
                        b = nil
                }
        }
	}
 
	func main() {
        a, b := make(chan bool), make(chan bool)
        t0 := time.Now()
        go func() {
                close(a)
                close(b)
        }()
        WaitMany(a, b)
        fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
	}

&emsp;&emsp;在重写的 WaitMany() 中，一旦接收到一个值，就将 a 或 b 的引用设置为 nil。当 nil channel 是 select 语句的一部分时，它实际上会被忽略，因此，将 a 设置为 nil 便会将其从 select 中移除，仅仅留下 b 等待它被关闭，进而退出循环。

&emsp;&emsp;总的来说，close和nil channel这些特性非常简单，但是它们的功能强大，使得创建高并发程序变得简单。

[返回顶部]({{ BASE_PATH }}{{ site.categories.usage.first.url }})


