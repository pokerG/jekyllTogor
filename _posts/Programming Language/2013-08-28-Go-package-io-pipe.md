---
date: 2013-08-28
layout:  post
title:  Go标准库-io-pipe
tagline:  ""
categories:
-  Programming Language
tags:
- Go
- Package
---

&emsp;&emsp;pipe中基本的struct是 pipe
	
	//pipe 是 PipeReader 和 PipeWriter 的底层实现
	type pipe struct {
		//sync.Mutex 是 一个并发时用的锁，这个以后在sync包中会讲到
		//rl/wl 控制同时只有一个读取器或写入器
		rl    sync.Mutex 
		wl    sync.Mutex 
		l     sync.Mutex //用于保护其他字段
		data  []byte     //管道中的数据
		// r/wwait 控制读取器或写入器等待	
		rwait sync.Cond  
		wwait sync.Cond  
		// r/werr 如果读取器（写入器）关闭，该错误会被Write（Read）方法返回
		rerr  error      
		werr  error     
	}
	
	//ErrClosePipe 用于返回对于关闭的管道的操作错误
	var ErrClosedPipe = errors.New("io: read/write on closed pipe")

	func (p *pipe) read(b []byte) (n int, err error) {
    // One reader at a time.（控制一次只能一个读取器）
    	p.rl.Lock()
    	defer p.rl.Unlock()

    	// 保护其他字段的读写
    	p.l.Lock()
    	defer p.l.Unlock()
    	for {
    	    // Reader端关闭后，再Read，则返回ErrClosedPipe
    	    if p.rerr != nil {
    	        return 0, ErrClosedPipe
    	    }
    	    // 管道中有数据，退出循环
    	    if p.data != nil {
    	        break
    	    }
    	    // Writer端关闭，返回p.werr
    	    if p.werr != nil {
    	        return 0, p.werr
    	    }
    	    // 没有数据或管道没有关闭，读取端等待
    	    p.rwait.Wait()
    	}
    	// 管道中有数据，将其copy一份到b中
    	n = copy(b, p.data)
    	p.data = p.data[n:]
    	// 如果管道数据被读光，需要唤醒在等待的Writer
    	if len(p.data) == 0 {
    	    p.data = nil
    	    p.wwait.Signal()
    	}
    	return
	}

	func (p *pipe) write(b []byte) (n int, err error) {
    // pipe uses nil to mean not available
    	if b == nil {
        	// zero的定义为：var zero [0]byte
        	b = zero[:]
    	}

    	// One writer at a time.
    	p.wl.Lock()
    	defer p.wl.Unlock()

    	p.l.Lock()
    	defer p.l.Unlock()
    	
    	p.data = b
    	// 唤醒在等待的Reader
    	p.rwait.Signal()
    	for {
    	    // 数据被读走，退出循环
    	    if p.data == nil {
    	        break
    	    }
    	    // Reader端关闭，设置err = p.rerr，退出循环
    	    if p.rerr != nil {
    	        err = p.rerr
    	        break
    	    }
    	    // Writer端关闭后，再Writer，设置err = 	、ErrClosedPipe
        	if p.werr != nil {
            	err = ErrClosedPipe
        	}
        	// 数据没被读走（全部）或管道读取端没关闭，则等待
        	p.wwait.Wait()
    	}
    	// 计算写入的字节数
    	n = len(b) - len(p.data)
    	p.data = nil // in case of rerr or werr
    	return
	}

	//关闭读取器
	func (p *pipe) rclose(err error) {
		if err == nil {
			err = ErrClosedPipe
		}
		p.l.Lock()
		defer p.l.Unlock()
		p.rerr = err
		p.rwait.Signal()
		p.wwait.Signal()
	}

	//关闭写入端
	func (p *pipe) wclose(err error) {
		if err == nil {
			err = EOF
		}
		p.l.Lock()
		defer p.l.Unlock()
		p.werr = err
		p.rwait.Signal()
		p.wwait.Signal()
	}

&emsp;&emsp;pipe中对外的两个结构是PipeReader 和 PipeWriter。<br/>
&emsp;&emsp;PipeReader 实现了io.Reader 和 io.Closer<br/>
&emsp;&emsp;PipeWriter 实现了io.Wrter 和 io.Closer<br/>
&emsp;&emsp;这两个结构的Read/Write 事实上是调用了pipe的实现。

&emsp;&emsp;关于 Read 方法的说明：从管道中读取数据。该方法会堵塞，直到管道写入端开始写入数据或写入端关闭了。如果写入端关闭时带上了error（即调用CloseWithError关闭），该方法返回的err就是写入端传递的error；否则err为EOF。

&emsp;&emsp;关于 Write 方法的说明：写数据到管道中。该方法会堵塞，直到管道读取端读完所有数据或读取端关闭了。如果读取端关闭时带上了error（即调用CloseWithError关闭），该方法返回的err就是读取端
传递的error；否则err为 ErrClosedPipe。

	//返回一组对应的管道输入输出
	func Pipe() (*PipeReader, *PipeWriter) {
		p := new(pipe)
		p.rwait.L = &p.l
		p.wwait.L = &p.l
		r := &PipeReader{p}
		w := &PipeWriter{p}
		return r, w
	}

&emsp;&emsp;pipe 用于在并发时实现输入输出操作，没有内部缓存，所以在并发调用时是安全的。


[返回顶部]({{ BASE_PATH }}{{ site.categories.usage.first.url }})



