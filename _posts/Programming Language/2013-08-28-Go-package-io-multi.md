---
date: 2013-08-28
layout:  post
title:  Go标准库-io-multi
tagline:  ""
categories:
-  Programming Language
tags:
- Go
- Package
---

	//接受多个Reader/Writer 
	//返回一个Reader/Writer
	//操作这个Reader/Writer 就相当于操作多个Reader/Writer
	func MultiReader(readers ...Reader) Reader {
		return &multiReader{readers}
	}
	
	func MultiWriter(writers ...Writer) Writer {
		return &multiWriter{writers}
	}

&emsp;&emsp;multi中定义了两个内部struct，multiReader 和 multiWriter 分别实现了 io.Reader 和 io.Writer

	type multiReader struct {
		readers []Reader
	}
	type multiWriter struct {
		writers []Writer
	}

	func (mr *multiReader) Read(p []byte) (n int, err error) {
		for len(mr.readers) > 0 {
			n, err = mr.readers[0].Read(p)
			if n > 0 || err != EOF {
				if err == EOF {
					// Don't return EOF yet. There may be more bytes
					// in the remaining readers.
					err = nil
				}
				return
			}
			mr.readers = mr.readers[1:]
		}
		return 0, EOF
	}

	func (t *multiWriter) Write(p []byte) (n int, err error) {
		for _, w := range t.writers {
			n, err = w.Write(p)
			if err != nil {
				return
			}
			if n != len(p) {
				err = ErrShortWrite
				return
			}
		}
		return len(p), nil
	}

&emsp;&emsp;我们举两个例子，来谈谈应用。

**MultiReader:**

	readers := []io.Reader{
	    strings.NewReader("from strings reader"),
	    bytes.NewBufferString("from bytes buffer"),
	}
	reader := io.MultiReader(readers...)
	data := make([]byte, 0, 1024)
	var (
	    err error
	    n   int
	)
	for err != io.EOF {
	    tmp := make([]byte, 512)
	    n, err = reader.Read(tmp)
	    if err == nil {
	        data = append(data, tmp[:n]...)
	    } else {
	        if err != io.EOF {
	            panic(err)
	        }
	    }
	}
	fmt.Printf("%s\n", data)

输出

	from strings readerfrom bytes buffer

&emsp;&emsp;MultiReader只是逻辑上将多个Reader组合起来，并不能通过调用一次Read方法获取所有Reader的内容。在所有的Reader内容都被读完后，Reader会返回EOF。

**MultiWriter：**

	file, err := os.Create("tmp.txt")
	if err != nil {
    	panic(err)
	}
	defer file.Close()
	writers := []io.Writer{
	    file,
	    os.Stdout,
	}
	writer := io.MultiWriter(writers...)
	writer.Write([]byte("MultiWriter test"))

&emsp;&emsp;这段程序执行后在生成tmp.txt文件，同时在文件和屏幕中都输出：MultiWriter test 这和Unix中的tee命令类似。
