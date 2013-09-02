---
date: 2013-08-26
layout:  post
title:  interface
tagline:  "Golang的哲学变革"
categories:
-  Programming Language
tags:
- Go
---

> If I could export one feature of Go into other languages, it would be interfaces.
> &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;———Rob Pike

>&emsp;&emsp;interface作为Golang的整个类型系统的基石，让Golang在基础编程哲学的探索上达到前所未有的高度。>&emsp;&emsp;&emsp;&emsp;——《Go语言编程》

&emsp;&emsp;接口是Java引入的特性。在软件工程上占有重要地位。而Golang变革了接口的概念。

&emsp;&emsp;我们考虑这样一个情况，有两个struct都实现了一些非常有用的小子集中的相关方法，这时有办法操作两个struct的任意一个就显得非常有用。那么我们该怎么办呢？这就是接口的作用。

	type s1 struct{
		...//一些属性
	}
	type s2 struct{
		...//一些属性
	｝
	func (a *s1)foo(){
		...//具体实现1
	}
	func (b *s2)foo(){
		...//具体实现2
	｝

&emsp;&emsp;这时，如果我们有需要操作两个stuct中的任意一个，可以声明这样一个interface

	type Foo interface{
		foo()
		//只要实现了foo()函数，就相当于实现了Foo接口 
		//我们就可以用接口来调用对象

	}
	
	func test(a Foo,b int){
		a.foo()
		//传进不同的struct对象，函数将调用不同的foo()方法
	}
	

&emsp;&emsp;从这个例子，我们就能看出interface的作用，降低耦合性，可以让某个模块或功能重复利用。

&emsp;&emsp;我们来看看io包，其中最基本的也是最重要的两个接口Reader和Writer。我们说，只要实现了这两个接口，就基本拥有了IO的功能。

	func ReadFrom(reader io.Reader,num int)([]byte,error){
		p := make([]byte,num)
		n,err：＝　reader.Read(p)
		if n > 0 {
 			return p[:n],nill
		}
		return p,err
	}

&emsp;&emsp;ReadFrom将io.Reader作为参数，我们可以从任何实现了io.Reader接口的地方读取数据

	//从标准输入流读取
	data,err := ReadFrom（os.Stdin,10)

	//从文件中读取
	file, err1 := os.Open(util.GetProjectRoot() + "01.txt")
	data,err2 := ReadFrom（file,10)

	//从字符串读取
	data，err := ReadFrom(strings.NewReader("from string"),10)

&emsp;&emsp;这是strings包中关于Reader的Read的实现

	func (r *Reader) Read(b []byte) (n int, err error) {
		if len(b) == 0 {
			return 0, nil
		}
		if r.i >= len(r.s) {
			return 0, io.EOF
		}
		n = copy(b, r.s[r.i:])
		r.i += n
		r.prevRune = -1
		return
	}

&emsp;&emsp;这里没有强制声明strings.Reader这个struct实现了哪个interface，但是它实现了Read方法，在Golang中也就实现了io.Reader这个接口。这种隐式的声明在Golang中术语叫做非侵入式接口。这样做的好处，我曾经在 **初识Go语言** 中提过。

&emsp;&emsp;其一，Go语言的标准库，再也不需要绘制类库的继承树图。

&emsp;&emsp;其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理。接口由使用方按需定义，而不用事前规划。

&emsp;&emsp;其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口。

&emsp;&emsp;Golang也可以通过接口的组合，实现一种抽向上的层次，

	type ReadWriter interface {
		Reader
		Writer
	}

	type ReadCloser interface {
		Reader
		Closer
	}

	type WriteCloser interface {
		Writer
		Closer
	}

	type ReadWriteCloser interface {
		Reader
		Writer
		Closer
	}

&emsp;&emsp;这几个接口都是Reader，Writer，Closer等基本接口的组合，可以说接口组合为Go程序建立起一个严密而有序的体系。

&emsp;&emsp;最后再谈谈一个有趣的interface ——interface｛｝
由于Go语言中任何对象实例都满足空接口interface{}，所以interface{}看起来像是可以指向任何对象的Any类型。
	
	var v1 interface{} = 1 // 将int类型赋值给interface{}
	var v2 interface{} = "abc" // 将string类型赋值给interface{}
	var v3 interface{} = &v2 // 将*interface{}类型赋值给interface{}
	var v4 interface{} = struct{ X int }{1}
	var v5 interface{} = &struct{ X int }{1}

&emsp;&emsp;当函数可以接受任意的对象实例时，我们会将其声明为interface{}，最典型的例子是标准库fmt中PrintXXX系列的函数，例如：
	
	func Printf(fmt string, args ...interface{})
	func Println(args ...interface{})

####总结

&emsp;&emsp;interface作为Golang类型系统的基石，为Golang工程形成了一个严密而清晰的体系，但其理念和使用却极为简洁，是Golang众多优秀特性中极为突出的一个，是我们学习Golang必须掌握的一点。


[返回顶部]({{ BASE_PATH }}{{ site.categories.usage.first.url }})
