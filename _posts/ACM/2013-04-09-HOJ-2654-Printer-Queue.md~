---
layout: post
title: HOJ 2654  Printer Queue
category : ACM
tagline: ""
tags : [数据结构]
---
<p>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	The only printer in the computer science students' union is experiencing an extremely heavy workload. Sometimes there are a hundred jobs in the printer queue and you may have to wait for hours to get a single page of output. Because some jobs are more important than others, the Hacker General has invented and implemented a simple priority system for the print job queue. Now, each job is assigned a priority between 1 and 9 (with 9 being the highest priority, and 1 being the lowest), and the printer operates as follows.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<img src="http://acm.hit.edu.cn/hoj/static/img/pic/26541.bmp" alt="" style="margin:10px; padding:0px" />
</p>
<ul style="margin:0px; padding:15px 0px 15px 30px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; font-size:16px; line-height:20px; background-color:rgb(247,247,247)">
	<li style="margin:0px; padding:0px">
		The first job J in queue is taken from the queue.
	</li>
	<li style="margin:0px; padding:0px">
		If there is some job in the queue with a higher priority than job J, thenmove J to the end of the queue without printing it.
	</li>
	<li style="margin:0px; padding:0px">
		Otherwise, print job J (and do not put it back in the queue).
	</li>
</ul>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	In this way, all those importantmuffin recipes that the Hacker General is printing get printed very quickly. Of course, those annoying term papers that others are printing may have to wait for quite some time to get printed, but that's life.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Your problem with the new policy is that it has become quite tricky to determine when your print job will actually be completed. You decide to write a program to figure this out. The program will be given the current queue (as a list of priorities) as well as the position of your job in the queue, and must then calculate how long it will take until your job is printed, assuming that no additional jobs will be added to the queue. To simplifymatters, we assume that printing a job always takes exactly one minute, and that adding and removing jobs from the queue is instantaneous.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<h3 style="margin:0px; padding:0px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; line-height:20px; background-color:rgb(247,247,247)">
	Input
</h3>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	One line with a positive integer: the number of test cases (at most 100). Then for each test case:
</p>
<ul style="margin:0px; padding:15px 0px 15px 30px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; font-size:16px; line-height:20px; background-color:rgb(247,247,247)">
	<li style="margin:0px; padding:0px">
		One line with two integers n and m, where n is the number of jobs in the queue (1 ≤ n ≤ 100) and m is the position of your job (0 ≤ m ≤ n - 1). The first position in the queue is number 0, the second is number 1, and so on.
	</li>
	<li style="margin:0px; padding:0px">
		One linewith n integers in the range 1 to 9, giving the priorities of the jobs in the queue. The first integer gives the priority of the first job, the second integer the priority of the second job, and so on.
	</li>
</ul>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<h3 style="margin:0px; padding:0px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; line-height:20px; background-color:rgb(247,247,247)">
	Output
</h3>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	For each test case, print one line with a single integer; the number of minutes until your job is completely printed, assuming that no additional print jobs will arrive.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<h3 style="margin:0px; padding:0px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; line-height:20px; background-color:rgb(247,247,247)">
	Sample Input
</h3>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<pre style="margin-top:0px; margin-bottom:0px; padding:5px; background-color:rgb(224,224,224); font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">3
1 0
5
4 2
1 2 3 4
6 0
1 1 9 1 1 1</pre>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<h3 style="margin:0px; padding:0px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; line-height:20px; background-color:rgb(247,247,247)">
	Sample Output
</h3>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<pre style="margin-top:0px; margin-bottom:0px; padding:5px; background-color:rgb(224,224,224); font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">1
2
5</pre>
<br />

<p>
	<span style="font-size:24px">题解：</span>
</p>
<p>
	<span style="font-size:24px">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;</span><span style="font-size:18px">模拟个队列</span>
</p>
<p>
	<span style="font-size:18px">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 要注意 当front 找到 rear 没有比它大的点 time++</span>
</p>
<p>
	<span style="font-size:18px">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;如果这个点是自己的那个点，直接输出。</span>
</p>
<p>
	<span style="font-size:18px">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; 另外如果front指向的是自己的点，不能打印的话，标记位置也得跟着变 m = rear - 1</span>
</p>
<p>
	<span style="font-size:18px"><br />
	</span>
</p>
<p>
	<span style="font-size:18px"></span>
</p>
<pre name="code" class="cpp">#include &lt;cstdio&gt;
#include &lt;cstdlib&gt;
#include &lt;cstring&gt;
int front,rear;
int q[10005];
int main(){
	int cases,time,n,m,ok;
	scanf(&quot;%d&quot;,&amp;cases);
	while(cases--){
		front = rear = 0;
		time = 0;
		scanf(&quot;%d%d&quot;,&amp;n,&amp;m);
		for(int i = 0; i &lt; n; i++){
			scanf(&quot;%d&quot;,&amp;q[i]);
			rear++;
		}
		while(front &lt;= rear){
			int i;
			for(i = front; i &lt; rear; i++)
				if(q[i] &gt; q[front]){
					q[rear++] = q[front];
					break;
				}
			if(i &gt;= rear){
				time += 1;
				if(front == m){
					printf(&quot;%d\n&quot;,time);
					break;
				}
				else
					front ++;
			}
			else if(front == m) 
				m = rear - 1;
			else 
				front++;
		}
	}
	return 0;
}

		
</pre>
<br />
<br />

<p>
</p>
<p>
	<span style="font-size:18px">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;</span>
</p>
<p>
	<br />
	
</p>