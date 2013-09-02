---
layout: post
title: 初识Go语言[1]
category : Programming Language
tagline: "我对Go语言语法的一些看法"
tags : [Go]
---
{% include JB/setup %}

&emsp;作为一个刚接触Go没几天的Gopher,没什么能力对其原理和高级用法做出分析和评价，姑且就对其语法形式谈谈我的看法，也算是对我这几天来的一个总结和回顾，姑妄言之，大家莫怪。

变量的声明
----
&emsp;首先最基本的声明方式	 `var a,b int`
和C语言相比 仅仅是把 类型 和 变量名称的顺序调换，但是我认为这样看起来比C语言可读性更高，而且var的引入使注意力更加集中（可能是前段时间php，js，python什么的看多了吧)

&emsp;再者 ` a,b := 1,2` 这种声明方式确实让我耳目一新，Go语言是类型安全的，但是这种写法让我有了动态语言的感觉，使用起来很方便，也很舒服。


类型
---

&emsp;Go与C相比类型多了很多，complex,error,map,interface,chan 等等，其中最让我感兴趣的是map 和 interface（chan 是并发的内容，暂不考虑), interface 一会专门讲，我们先谈谈map，php等语言的数组就具有map类似的功能，Go的map使用起来也很方便，赋值，删除，查找做的非常简单和精巧。

&emsp;slice看起来很像指向数组的指针，但还有些许不同，它还有元素个数和分配的空间两个“属性”，slice和数组很容易让人联想到STL里的vector和数组的关系，slice的动态扩展空间的能力和各种管理功能很好很强大。

流程控制
---

&emsp;Go语言的流程控制语句少了小括号，现在很多语言都去掉了这个，让我们可以少打很多字

&emsp;switch语句默认加了break，这个确实让我很兴奋，不用再担心少写break而出现bug

&emsp;Go语言的循环语言只有for关键字（确实这是最常用的，C语言中我也很少用while和do-while），Go中for的形式非常灵活，功能也很完善。

&emsp;goto被许许多多的人所诟病，但是对于Go这样一个仅有25个关键字的语言来说，仍然支持goto关键字，说明goto在某些场合下确实是很合适的。（linux内核中goto出现的频率之高！！！）。break和continue 可以带标签，着实让我感慨，真是为程序员考虑啊。以后再也不用为多重循环的跳转头疼了。

函数
---

&emsp;Go语言函数的关键字，让我看着不太爽，func感觉像是被剪掉尾巴的孔雀，缩写的太难看，要不function 要不 fn 多好。

&emsp;因为Go的函数是多返回值的，所以它的函数头，返回值位于后头，这样写也挺好 `func add(a , b int)(ret int,err error)｛｝`

&emsp;Go函数可以有多个返回值，这个是迄今为止Go语言与其他语言最大的不同，结合上Go语言的多重赋值，Go语言的代码会很简洁、优雅。

######结尾

&emsp;今天我们就先谈论到这，明天我们说说Go语言的interface，error，和 OOP特性。

 [返回顶部]({{ BASE_PATH }}{{ site.categories.usage.first.url }})


