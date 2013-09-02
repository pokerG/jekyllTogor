---
layout: post
title: HOJ 2196 Job Scheduling by Open Bidding
category : ACM
tagline: ""
tags : [DP]
---
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Your team is setting up a computing resource devoted to batch processing of compute-bound jobs. In addition, you have decided to use static scheduling for each period of time. Naturally, you wish to maximize the income for each set of jobs run, and you have been given the responsibility of finding an optimal mix of jobs for each set of candidate jobs. The jobs are submitted by an open bid process: clients will specify the amount of processor time they wish to reserve and the dollar amount that they wish to pay. If a job finishes early, the client will still pay the full amount, and if a job exceeds the requested time, it will be terminated and (of course) the client will still pay the full amount. For purposes of scheduling, your team assumes that each job will in fact use its entire scheduled time slot. In the interests of good customer relations, though, you are not to include a bid in the schedule if there is not sufficient time available to satisfy it — we’re not going to over-book like the airlines do, and then hope someone doesn’t use the full allotment!
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Input</span>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	The input file begins with a line containing a single integer (no white space) specifying the number of problem sets in the file.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Each problem set consists of (n+2) lines (no white space except as specified):
</p>
<ul style="margin:0px; padding:15px 0px 15px 30px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; font-size:16px; line-height:20px; background-color:rgb(247,247,247)">
	<li style="margin:0px; padding:0px">
		a single integer n (n &lt;= 500) specifying the number of candidate jobs to be schedu
	</li>
	<li style="margin:0px; padding:0px">
		n lines giving the bid as an integer specifying the number of seconds followed by a single space and then a dollar amount given in decimal form (always showing two digits to the right of the decimal point)
	</li>
	<li style="margin:0px; padding:0px">
		a single integer t (t &lt;= 2000) specifying the amount of time to be scheduled with these job
	</li>
</ul>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Output</span>
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	Each problem set will be numbered (beginning at one) and will generate a single line: Problem&nbsp;:&nbsp;seconds scheduled for $abc.de where&nbsp;is replaced by the problem set number,&nbsp;is replaced with the total time actually scheduled (possibly not the full input time), and $abc.de is replaced by the dollar amount, given always with the leading currency symbol and with two digits to the right of the decimal point. There will be no blank lines, and the final line will end with the new-line character.
</p>
<p style="margin-top:0px; margin-bottom:5px; padding-top:0px; padding-bottom:0px; line-height:20px; font-size:14px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">
	<span style="margin:0px; padding:0px; font-size:18px; line-height:28px">Sample Input</span>
</p>
<pre style="margin-top:0px; margin-bottom:0px; background-color:rgb(224,224,224); padding:5px; font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">1
10
19 0.78
12 0.31
17 0.77
22 0.77
8 0.56
10 0.33
17 0.35
24 0.12
22 0.70
5 0.52
120</pre>
<span style="margin:0px; padding:0px; font-size:18px; line-height:28px; color:rgb(51,51,51); font-family:'Helvetica Neue',Helvetica,Arial,sans-serif; background-color:rgb(247,247,247)">Sample Output</span>
<pre style="margin-top:0px; margin-bottom:0px; background-color:rgb(224,224,224); padding:5px; font-size:14px; line-height:16px; overflow:auto; font-family:Consolas,'Lucida Console','Andale Mono','Bitstream Vera Sans Mono','Courier New',Courier; color:rgb(51,51,51)">Problem 1: 120 seconds scheduled for $4.78</pre>
<p>
	<strong><span style="font-size:32px">&nbsp;题解：&nbsp;</span></strong>
</p>
<p>
	<strong><span style="font-size:32px">&nbsp; &nbsp; &nbsp; &nbsp; &nbsp;&nbsp;</span><span style="font-size:24px">基本的01背包问题，直接给代码吧</span></strong>
</p>
<p>
	<strong></strong>
</p>
<pre name="code" class="html"><span style="font-size: 18px;">#include &lt;cstdio&gt;
#include &lt;cstdlib&gt;
#include &lt;cstring&gt;
int c[505];
double w[505];
double f[2005];
double w_max(double a,double b){
	return a &gt; b ? a : b;
}
int main(){
	int cases,n,r = 0;
	scanf(&quot;%d&quot;,&amp;cases);
	while(cases--){
		r++;
		memset(f,0,sizeof(f));
		scanf(&quot;%d&quot;,&amp;n);
		for(int i = 0; i &lt; n; i++)
			scanf(&quot;%d %lf&quot;,&amp;c[i],&amp;w[i]);
		int time;
		scanf(&quot;%d&quot;,&amp;time);
		for(int i = 0; i &lt; n; i++)
			for(int j = time; j &gt;= c[i]; j--)
				f[j] = w_max(f[j],f[j - c[i]] + w[i]);
		int actTime = 0;
		double maxSum = 0;
		for(int i = 1; i &lt;= time; i ++)
			if(f[i] &gt; maxSum)
				actTime = i, maxSum = f[i];
		printf(&quot;Problem %d: %d seconds scheduled for $%.2lf\n&quot;,r,actTime,maxSum);
	}
	return 0;
}</span><span style=" font-size: 24px;"><strong>
</strong></span></pre>