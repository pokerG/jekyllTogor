---
date: 2013-08-27
layout:  post
title:  Go标准库-io-io
tagline:  ""
categories:
-  Programming Language
tags:
- Go
- Package
---

&emsp;&emsp;io包为I/O原语提供了基本的接口。它主要包装了这些原语的已有实现。由于这些接口和原语以不同的实现包装了低级操作，因此除非另行通知，否则客户端不应假定它们对于并行执行是安全的。	

&emsp;&emsp;首先我们来看几个variables，这几个variables都是读写操作中可能出现的情况
	
	//ErrShortWrite 用于一个write少于规定的字符数，但是又不能返回一个明确的错误时
	var ErrShortWrite = errors.New("short write")

	//ErrShortBuffer 用于一个read要求的字符数多于所能提供的字符
	var ErrShortBuffer = errors.New("short buffer")

	//EOF 就是end-of-file 不用多解释，但是如果EOF发生在一个结构数据流中，error就不应该是EOF
	var EOF = errors.New("EOF")

	//ErrUnexpectedEOF 意味着EOF发生在读取一个固定数据块或者数据结构的中间
	var ErrUnexpectedEOF = errors.New("unexpected EOF")

	//ErrNoProgress 发生在多个io.Reader调用Read()方法，却没有返回任何数据或错误时
	//通常作为一个破损的io.Reader实现的标志
	var ErrNoProgress = errors.New("multiple Read calls return no data or error")

&emsp;&emsp;io包中最基本的4个接口

	//Reader 封装了基本的read方法
	//Read()方法读取len(p)个字符放入p中，返回读取的字符数（0<=n<=len(p)) 和 错误。
	//即使 Read 返回的 n < len(p)，它也会在调用过程中使用 p 的全部作为暂存空间。
	//若一些数据可用但不到 len(p) 个字节，Read 会照例返回可用的数据，而不是等待更多数据。
	//当 Read 在成功读取 n > 0 个字节后遇到一个错误或EOF（end-of-file），它就会返回读取的字节数。
	//所以当Read方法返回错误时，不代表没有读取到任何数据。调用者应该处理返回的任何数据，之后才处理可能的错误。
	//"p被装满" err返回nil
	//"数据被读完", err返回EOF
	//"读取出错", err返回相应的错误信息
	//以上就是实现io.Reader接口应注意的
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	//Writer 封装了基本的write方法
	//Write()方法 将p中数据写入相应对象的数据流，返回从p中写入的字符数和错误。
	//当n = len(p)时，返回nil
	//当 n < len(p)时，返回相应的错误信息
	type Writer interface {
		Write(p []byte) (n int, err error)
	}

	//Closer 封装了基本的Close方法
	//Close()用于关闭数据流，文件，数据库，Socket等等
	//经常将Close()的调用放在defer语句
	type Closer interface {
		Close() error
	}

	//Seeker 封装了基本的Seek方法
	//Seek() 设置下一次的读写头，即偏移地址offset
	//offset的解释依据于 whence
	// whence = 0 与文件起始的偏移
	// whence = 1 与当前读写头的偏移
	// whence = 2 与文件尾部的偏移
	type Seeker interface {
		Seek(offset int64, whence int) (ret int64, err error)
	}

&emsp;&emsp;其他接口都是这四个基本接口的组合，从名字我们就可以判断出来。
**ReadWriter,ReadCloser,WriteCloser,ReadWriteCloser,ReadSeeker,WriteSeeker,ReadWriteSeeker**

&emsp;&emsp;这里还有一些较为高级一点的输入输出接口

	//ReaderFrom 封装了ReadFrom方法
	//ReadFrom() 从 r 中读取数据，直到EOF或 error
	//返回 读取的字符数 和 error
	//Copy() 如果相应的参数实现了ReaderFrom接口，会调用ReadFrom()
	type ReaderFrom interface {
		ReadFrom(r Reader) (n int64, err error)
	}

	//WriterTo 封装了WriterTo方法
	//WriteTo()将对象的数据流写入w 直到全部写入或者遇到error
	//返回 写入的字符数 和 error
	//Copy() 如果相应的参数实现了WriterTo接口，会调用WriteTo()
	type WriterTo interface {
		WriteTo(w Writer) (n int64, err error)
	}

	//ReaderAt 封装了ReadAt方法
	//ReadAt() 从对象的数据流的off位置读取len(p)个字符 到 p
	//返回 读取的字符数 和 error
	//当 ReadAt 返回的 n < len(p) 时，它就会返回一个非nil的错误来解释 
	为什么没有返回更多的字节。在这一点上，ReadAt 比 Read 更严格。
	//即使 ReadAt 返回的 n < len(p)，它也会在调用过程中使用
	 p 的全部作为暂存空间。若一些数据可用但不到 len(p) 字节，
	ReadAt 就会阻塞直到所有数据都可用或产生一个错误。 
	在这一点上 ReadAt 不同于 Read。
	// 返回读取的字节数 n 和读取时遇到的错误
	//如果 n < len(p)，则需要返回一个 err 值来说明
	// 为什么没有将 p 填满（比如 EOF）
	// 如果 n = len(p)，而且对象的数据没有全部读完，则
	// err 将返回 nil
	// 如果 n = len(p)，而且对象的数据刚好全部读完，则
	// err 将返回 EOF 或者 nil（不确定）
	type ReaderAt interface {
		ReadAt(p []byte, off int64) (n int, err error)
	}

	//WriterAt 封装了WriterAt方法
	//WriteAt() 将 p 中的数据写入到对象数据流的 off 处
	//返回 写入的字符数 和 error
	type WriterAt interface {
		WriteAt(p []byte, off int64) (n int, err error)
	}

	//ByteReader 封装了ReadByte方法
	//ReadByte 从对象的数据流中读取一个字符到 c
	//如何没有字符可读，返回一个error
	type ByteReader interface {
		ReadByte() (c byte, err error)
	}

	//ByteScanner 在ByteReader的基础上加上了UnreadByte方法
	//UnreadByte() 撤销上一次的ReadByte 即将读写头移到上次ReadByte的位置
	//调用UnreadByte()前必须调用ReadByte(),且不能连续调用UnreadByte()两次，否则返回一个错误
	type ByteScanner interface {
		ByteReader
		UnreadByte() error
	}

	//ByteWriter 封装了 WriteByte方法
	//WriteByte 将一个字节 c 写入到对象的数据流中
	type ByteWriter interface {
		WriteByte(c byte) error
	}

	//RuneReader 封装了ReadRune方法
	//ReadRune() 从对象的数据流中读取一个UTF-8 字符 
	//返回这个 rune, 还有它所占字符数，还有error
	type RuneReader interface {
		ReadRune() (r rune, size int, err error)
	}

	//RuneScanner 类似于 ByteScanner 就不再赘述
	type RuneScanner interface {
		RuneReader
		UnreadRune() error
	}

	//stringWriter 封装了 WriteString方法
	//WriteString() 将 字符串s 写入 对象的数据流
	//返回写入的字符数 和 error
	type stringWriter interface {
		WriteString(s string) (n int, err error)
	}

&emsp;&emsp;接下来看看io中的几个函数

	//向 Writer 中 写入 字符串
	//返回 写入的字符书 和 error
	//如果 w 也实现了 stringWriter接口 将直接调用其WriteString()
	func WriteString(w Writer, s string) (n int, err error) {
		if sw, ok := w.(stringWriter); ok {
			return sw.WriteString(s)
		}
		return w.Write([]byte(s))
	}
	
	//从Reader中读取至少min个字符 到 buf
	//返回 读取的字符数 和 error
	//只有当没有数据可读，返回EOF
	//如果 n < min 返回ErrUnexpectedEOF
	//如果 min > len(buf) 返回 ErrShortBuffer
	//只有当 n >= min 返回nil
	func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
		if len(buf) < min {
			return 0, ErrShortBuffer
		}
		for n < min && err == nil {
			var nn int
			nn, err = r.Read(buf[n:])
			n += nn
		}
		if n >= min {
			err = nil
		} else if n > 0 && err == EOF {
		err = ErrUnexpectedEOF
		}
		return
	}

	//ReadFull 与 ReadAtLeast相似，min = len(buf)
	func ReadFull(r Reader, buf []byte) (n int, err error) {
		return ReadAtLeast(r, buf, len(buf))
	}

	//从 Reader 中 复制 n个字符 到 Writer
	//返回 复制的字符数 和 复制过程中最早遇到的错误
	//只有 written == n 返回 nil
	//
	func CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
		written, err = Copy(dst, LimitReader(src, n))
		if written == n {
			return n, nil
		}
		if written < n && err == nil {
			// src stopped early; must have been EOF.
			err = EOF
		}
		return
	}

	//将 src 复制到 dst，直到在 src 上到达 EOF 或发生错误。
	//返回 复制的字符数 和 复制过程中最早遇到的错误
	//成功的 Copy 返回 err == nil，而非 err == EOF。
	//由于 Copy 被定义为从 src 读取直到 EOF 为止，
	//因此它不会将来自 Read 的 EOF 当做错误来报告。
	func Copy(dst Writer, src Reader) (written int64, err error) {

		if rt, ok := dst.(ReaderFrom); ok {
			return rt.ReadFrom(src)
		}
	
		if wt, ok := src.(WriterTo); ok {
			return wt.WriteTo(dst)
		}
		buf := make([]byte, 32*1024)
		for {
			nr, er := src.Read(buf)
			if nr > 0 {
				nw, ew := dst.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = ErrShortWrite
					break
				}
				}
			if er == EOF {
				break
			}
			if er != nil {
				err = er
				break
			}
		}
		return written, err
	}

&emsp;&emsp;最后看看io中定义的三个struct

LimiteReader结构

	//LimiteReader 从R 中 读取　但限制为Ｎ个字符
	//每次调用Read() 都将更新N
	type LimitedReader struct {
		R Reader // underlying reader
		N int64  // max bytes remaining
	}
	
	//LimitReader 用来生成 LimiteReader实例
	func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

	func (l *LimitedReader) Read(p []byte) (n int, err error) {
		if l.N <= 0 {
			return 0, EOF //读取完毕后 返回EOF
		}
		if int64(len(p)) > l.N {
			p = p[0:l.N]
		}
		n, err = l.R.Read(p)
		l.N -= int64(n) 	//更新N 来表示 剩余能读取的字符数
		return
	}

SectionReader结构

	//SectionReader是一个struct（没有任何导出的字段）
	//实现了 Read, Seek 和 ReadAt，同时，内嵌了 ReaderAt 接口。
	type SectionReader struct {
		r     ReaderAt 
		base  int64 //基址 NewSectionReader 会将 base 设置为 off
		off   int64 //从 r 中的 off 偏移处开始读取数据
		limit int64 // limit - off = SectionReader 流的长度
	}
	
	// 返回一个 SectionReader
	//它从 r 中的偏移量 off 处读取 n 个字节后以 EOF 停止
	func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader {
		return &SectionReader{r, off, off, off + n}
	}

	func (s *SectionReader) Read(p []byte) (n int, err error) {
		if s.off >= s.limit {
			return 0, EOF
		}
		if max := s.limit - s.off; int64(len(p)) > max {
			p = p[0:max]
		}
		n, err = s.r.ReadAt(p, s.off)
		s.off += int64(n)
		return
	}
	
	//Seek() 中的 两种错误
	var errWhence = errors.New("Seek: invalid whence")
	var errOffset = errors.New("Seek: invalid offset")
	
	func (s *SectionReader) Seek(offset int64, whence int) (ret int64, err error) {
		switch whence {
		default:
			return 0, errWhence
		case 0:
			offset += s.base
		case 1:
			offset += s.off
		case 2:
			offset += s.limit
		}
		if offset < s.base || offset > s.limit {
			return 0, errOffset
		}
		s.off = offset
		return offset - s.base, nil
	}

	func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error) {
		if off < 0 || off >= s.limit-s.base {
			return 0, EOF
		}
		off += s.base
		if max := s.limit - off; int64(len(p)) > max {
			p = p[0:max]
			n, err = s.r.ReadAt(p, off)
			if err == nil {
				err = EOF
			}
			return n, err
		}
		return s.r.ReadAt(p, off)
	}

	// 返回 SectionReader的大小
	func (s *SectionReader) Size() int64 { 
		return s.limit - s.base 
	}

teeReader结构

	//将从 r 中读到的数据写入 w 中。
	//所有经由它处理的从 r 的读取都匹配于对应的对 w 的写入。
	//它没有内部缓存，即写入必须在读取完成前完成。
	//任何在写入时遇到的错误都将作为读取错误返回。
	type teeReader struct {
		r Reader
		w Writer
	}

	func TeeReader(r Reader, w Writer) Reader {
		return &teeReader{r, w}
	}

	func (t *teeReader) Read(p []byte) (n int, err error) {
		n, err = t.r.Read(p)
		if n > 0 {
			if n, err := t.w.Write(p[:n]); err != nil {
				return n, err
			}
		}
		return
	}

[返回顶部]({{ BASE_PATH }}{{ site.categories.usage.first.url }})
