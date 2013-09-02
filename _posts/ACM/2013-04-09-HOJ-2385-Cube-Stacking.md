---
date: 2013-04-09
layout:  post
title:  HOJ 2385 Cube Stacking
tagline:  ""
categories:
-  ACM
tags:
- 数据结构
---

<p>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Farmer John and Betsy are playing a game with N (1 &lt;= N &lt;= 30,000)identical cubes labeled 1 through N. They start with N stacks, each containing a single cube. Farmer John asks Betsy to perform P (1&lt;= P &lt;= 100,000) operation. There are two types of operations: moves and counts.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<ul style="margin:0px; padding:15px 0px 15px 30px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; font-size:16px; line-height:20px; background-color:rgb(247,247,247)">
	<li style="margin:0px; padding:0px">
		In a move operation, Farmer John asks Bessie to move the stack containing cube X on top of the stack containing cube Y.
	</li>
	<li style="margin:0px; padding:0px">
		In a count operation, Farmer John asks Bessie to count the number of cubes on the stack with cube X that are under the cube X and report that value.
	</li>
</ul>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Write a program that can verify the results of the game.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Input</span>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	* Line 1: A single integer, P
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	* Lines 2..P+1: Each of these lines describes a legal operation. Line 2 describes the first operation, etc. Each line begins with a 'M' for a move operation or a 'C' for a count operation. For move operations, the line also contains two integers: X and Y.For count operations, the line also contains a single integer: X.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Note that the value for N does not appear in the input file. No move operation will request a move a stack onto itself.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Output</span>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Print the output from each of the count operations in the same order as the input file.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Sample Input</span>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<pre style="margin-top:0px; margin-bottom:0px; padding:5px; background-color:rgb(224,224,224); font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">6
M 1 6
C 1
M 2 4
M 2 6
C 3
C 4</pre>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Sample Output</span>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
</p>
<pre style="margin-top:0px; margin-bottom:0px; padding:5px; background-color:rgb(224,224,224); font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">1
0
2</pre>
<br />

<p>
	<span style="font-size:24px">题解：</span>
</p>
<p>
	&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; <span style="font-size:18px">并查集，并记录i元素上面和下面的元素个数</span>
</p>
<p>
</p>
<pre name="code" class="cpp">#include &lt;cstdio&gt;
#include &lt;cstdlib&gt;
#include &lt;cstring&gt;
const int N = 30005;
int set[N];
int high[N]; //i元素上面（加上i）的元素个数 
int low[N];  //i元素下面的元素个数 
int findset(int x){
	if(x == set[x])
		return set[x];
	int tm = set[x];
	set[x] = findset(set[x]);
	low[x] += low[tm];
	return set[x];
}
void unionset(int x,int y){
	int fx = findset(x);
	int fy = findset(y);
	if( fx == fy)
		return ;
	set[fx] = fy;
	low[fx] += high[fy];
	high[fy] += high[fx];
}
int main(){
	int n;
	while(scanf(&quot;%d&quot;,&amp;n) != EOF){
		for(int i = 1; i &lt;= N; i++){
			set[i] = i;
		         high[i] = 1;
		}
		memset(low,0,sizeof(low));
		char op;
		int x,y;
		getchar();
		for(int i =1 ; i &lt;= n; i++){
			scanf(&quot;%c&quot;,&amp;op);
			if(op == 'C'){
				scanf(&quot;%d&quot;,&amp;x);
				findset(x);
				printf(&quot;%d\n&quot;,low[x]);
			}
			else{
				scanf(&quot;%d%d&quot;,&amp;x,&amp;y);
				unionset(x,y);
			}
			getchar();
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
	<br />
	
</p>
