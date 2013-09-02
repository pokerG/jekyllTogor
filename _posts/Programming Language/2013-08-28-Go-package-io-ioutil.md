---
date: 2013-08-28
layout:  post
title:  Go标准库-io-ioutil
tagline:  ""
categories:
-  Programming Language
tags:
- Go
- Package
---

&emsp;&emsp;ioutil封装了一些常用的io操作

	//读取r中所有数据
	//返回 读取的数据 和 error
	//读取成功的话 返回nil 而不是 EOF
	//ReadAll 实际上是通过bytes.Buffer的ReadFrom()实现的
	func ReadAll(r io.Reader) ([]byte, error)

	//读取文件中的所有数据
	//返回 读取的数据 和 error
	//读取成功的话 返回nil 而不是 EOF
	//ReadFile 和 ReadAll 都是通过readAll()实现的
	//ReadFile会先判断文件的大小，给bytes.Buffer一个预定义容量，避免额外分配内存。
	func ReadFile(filename string) ([]byte, error)

	//将 data 写入 文件
	//如果文件不存在，将以perm权限创建文件
	func WriteFile(filename string, data []byte, perm os.FileMode) error

    //定义了一系列文件信息的操作
	type byName []os.FileInfo

	func (f byName) Len() int           { return len(f) }
	func (f byName) Less(i, j int) bool { return f[i].Name() < f[j].Name() }
	func (f byName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

	//读取文件夹中的文件
	//返回 文件信息的序列 和 error
	func ReadDir(dirname string) ([]os.FileInfo, error)

&emsp;&emsp;这里有个实现类似于tree命令的程序

	func main() {
	    if len(os.Args) > 1 {
	        Tree(os.Args[1], 1, map[int]bool{1:true})
	    }
	}

	// 列出dirname目录中的目录树，实现类似Unix中的tree命令
	// curHier 当前层级（dirname为第一层）
	// hierMap 当前层级的上几层是否需要'|'的映射
	func Tree(dirname string, curHier int, hierMap map[int]bool) error {
	    dirAbs, err := filepath.Abs(dirname)
	    if err != nil {
	        return err
	    }
	    fileInfos, err := ioutil.ReadDir(dirAbs)
	    if err != nil {
	        return err
	    }
	
	    fileNum := len(fileInfos)
	    for i, fileInfo := range fileInfos {
	        for j := 1; j < curHier; j++ {
	            if hierMap[j] {
	                fmt.Print("|")
	            } else {
	                fmt.Print(" ")
	            }
	            fmt.Print(strings.Repeat(" ", 3))
	        }
	
	        // map是引用类型，所以新建一个map
	        tmpMap := map[int]bool{}
	        for k, v := range hierMap {
	            tmpMap[k] = v
	        }
	        if i+1 == fileNum {
	            fmt.Print("`")
	            delete(tmpMap, curHier)
	        } else {
	            fmt.Print("|")
	            tmpMap[curHier] = true
	        }
	        fmt.Print("-- ")
	        fmt.Println(fileInfo.Name())
	        if fileInfo.IsDir() {
	            Tree(filepath.Join(dirAbs, fileInfo.Name()), curHier+1, tmpMap)
        	}
		}
    	return nil
	}
&emsp;&emsp;

	//将 io.Reader 封装成 io.ReadCloser
	// 其 Close 方法不做任何事情
	func NopCloser(r io.Reader) io.ReadCloser

	//Discard 对应的类型 type devNull int 实现了io.Writer
	//从名字我们就可以看出 Discard	的 任何Write()都会成功但是不干任何事情
	//同时为了优化io.Copy到Discard，避免不必要的工作，实现了io.ReaderFrom接口。
	//ReadFrom的实现是读取内容到一个buf中，最大也就8192字节，其他的会丢弃（当然，这个也不会读取）
	//Discard 用于并发
	var Discard io.Writer = devNull(0)

&emsp;&emsp; tempfile.go中的函数

	// TempFile 在目录 dir 中创建一个临时文件并将其打开
	// 文件名以 prefix 为前缀
	// 返回创建的文件的对象和创建过程中遇到的任何错误
	// 如果 dir 为空，则在系统的临时目录中创建临时文件
	// 如果环境变量中没有设置系统临时目录，则在 /tmp 中创建临时文件
	// 调用者可以通过 f.Name() 方法获取临时文件的完整路径
	// 调用 TempFile 所创建的临时文件，应该由调用者自己移除
	func TempFile(dir, prefix string) (f *os.File, err error)

	// TempDir 功能同 TempFile，只不过创建的是目录
	// 返回值也只返目录的完整路径
	func TempDir(dir, prefix string) (name string, err error)

