---
date: 2013-08-20
layout:  post
title:  Go Build in
tagline:  "GO语言的内置关键字，类型及函数"
categories:
-  Programming Language
tags:
- Go
---

内置关键字（25个 均为小写）
-------
>break&emsp;&emsp;default&emsp;&emsp;func&emsp;&emsp;interface&emsp;&emsp;select<br />
>case&emsp;&emsp;defer&emsp;&emsp;go&emsp;&emsp;map&emsp;&emsp;struct<br />
>chan&emsp;&emsp;else&emsp;&emsp;goto&emsp;&emsp;package&emsp;&emsp;switch<br />
>const&emsp;&emsp;fallthrough&emsp;&emsp;if&emsp;&emsp;range&emsp;&emsp;type<br />
>continue&emsp;&emsp;for&emsp;&emsp; import&emsp;&emsp;return&emsp;&emsp;var <br/>

内置类型
------
>值类型：<br />
>bool<br />
>byte<br />
>int(32 or 64), int8, int16, int32, int64<br />
>uint(32 or 64), uint8(byte), uint16, uint32, uint64<br />
>float32, float64<br />
>string<br />
>complex64, complex128   --复数<br />
>array    -- 固定长度的数组<br />

>引用类型：(指针类型)<br />
>slice   -- 切片<br />
>map        -- 映射<br />
>chan    -- 管道<br />

内置函数
----
>append  -- 把东西增加到slice里面,返回修改后的slice<br />
>close   -- 关闭channel<br />
>delete    -- 从map中删除key对应的value<br />
>panic    -- 停止常规的goroutine<br />
>recover -- 允许程序定义goroutine的panic动作<br />
>imag    -- 返回complex的实部<br />
>real    -- 返回complex的虚部<br />
>make    -- 返回Type本身(只能应用于slice, map, channel)<br />
>new        -- 返回指向Type的指针<br />
>cap        -- 容量，容积capacity<br />
>copy    -- 复制slice，返回复制的数目<br />
>len        -- 返回长度<br />

内置接口
----
	type error interface {        //只要实现了Error()函数，返回值为String的都实现了err接口
		Error() string
	}

