---
layout: post
title: HOJ 1019 Grandpa's Other Estate
category : ACM
tagline: ""
tags : [模拟]
---
From our previous contest, we know that Kamran the Believer inherited many of his grandpa's belongings. Apparently, his grandpa had been a mathematician in his life with interests in puzzle solving, since he has made Kamran solve another programming problem!<br />

<p>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Grandpa had a big garden with many valuable walnut trees. He has written in his will that Kamran can inherit one piece of square shaped land of a given size in the garden, such that its sides be parallel to the x and y axes. Taking advantage of the fact that no other restrictions have been mentioned in the will, Kamran wants to choose the land in which the most number of trees lie. Kamran is too wealthy now and thus too lazy to spend time and solve another algorithmic problem. He has hired you to solve this problem for him.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	You are given the location of all trees in the big garden and the size of the land to choose. You are to write a program to find out where to choose the land so that the most number of trees lie in it. You may consider trees as points in the plane and the land as a square. You are to find the position of the square such that it includes as many points as possible. Note that the points on the border of the square are considered to be inside it.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Input</span><br style="margin:0px; padding:0px" />
	The first line of the input file contains a single integer t (1&lt;=t&lt;=10), the number of test cases, followed by the input data for each test case. The first line of each test case contains an integer n (1&lt;=n&lt;=100), the number of trees, and an integer r (1&lt;=r&lt;=1000), the length of the land's side, followed by n lines, each containing two integers x and y (0&lt;=x , y &lt;= 100,000) representing the coordinates of a walnut tree. Note that all coordinates are pairwise distinct.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Output</span><br style="margin:0px; padding:0px" />
	There should be one line per test case containing the maximum number of trees that Kamran can own.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Sample Input</span>
</p>
<pre style="margin-top:0px; margin-bottom:0px; padding:5px; background-color:rgb(224,224,224); font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">1
3 1
1 2
2 1
4 3</pre>
<p>
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">Sample Output</span>
</p>
<p>
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)"></span>
</p>
<pre style="margin-top:0px; margin-bottom:0px; background-color:rgb(224,224,224); padding:5px; font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">2</pre>
<br />

<p>
</p>
<p>
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)"><strong>题解：</strong></span>
</p>
<p>
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)"><strong>&nbsp; &nbsp; &nbsp;只需将正方形从左到右，从上到下扫描一遍，更新res, 即可</strong></span>
</p>
<p>
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)"></span>
</p>
<pre name="code" class="cpp">#include &lt;cstdio&gt;
#include &lt;cstdlib&gt;
#include &lt;algorithm&gt;
#include &lt;cmath&gt;
using namespace std;
struct point
{
	int x,y;
};
point p[105];
bool judge(int x1,int x2,int y2,int y1,point a){
	if((a.x - x1) * (a.x - x2) &lt;= 0 &amp;&amp; (a.y - y1) * (a.y - y2) &lt;= 0)
		return true;
	return false;
}
int main(){
	int t,n,r;
	scanf(&quot;%d&quot;,&amp;t);
	while(t--){
		scanf(&quot;%d%d&quot;,&amp;n,&amp;r);
		for(int i = 1; i &lt;= n; i++)
			scanf(&quot;%d%d&quot;,&amp;p[i].x,&amp;p[i].y);
		int res = 1;
		int temp;
		for(int i = 1; i &lt;= n; i ++)
			for(int j = i + 1; j &lt;= n; j ++){
				if (abs(p[i].x - p[j].x) &gt; r || abs(p[i].y - p[j].y) &gt; r)continue;
				int tx = min(p[i].x,p[j].x);
				int ty = max(p[i].y,p[j].y);
				temp = 0;
				for(int k = 1; k &lt;= n; k++)
					if(judge(tx,tx + r, ty , ty - r,p[k]))
						temp ++;
				res = max(res,temp);
			}
		printf(&quot;%d\n&quot;,res);
	}
	return 0;
}</pre>
<br />
<br />

<p>
</p>
<p>
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)"><br />
	</span>
</p>