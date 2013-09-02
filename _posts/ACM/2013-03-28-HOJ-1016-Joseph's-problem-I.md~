---
layout: post
title: 转： HOJ 1016 Joseph's problem I
category : ACM
tagline: ""
tags : [模拟]
---
<div style="">
	<div id="sina_keyword_ad_area2" style="line-height:25px">
		<div style="line-height:25px">
			先说说经典的Joseph问题吧，说有n个要被处决的人（编号0~(n-1))，从0开始报数，报到(m-1)的会被杀掉，剩下的人继续从0开始报数，如此下去最后剩的一个人会存活下来。说Joseph发现了这个规律而且把他透露了出来，现在假如你在这n个人里面，你会选择几号位置站下。
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　很显然你会选择能活下来的那个位置，所以问题就是如何得到这个位置。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　首先想到的是模拟(至少我笨脑子是这么想的)，但无论是用链表还是用数组这个时间复杂度都是比较高的，至少交题的时候会TLE，这里介绍一种线性时间的解法，出自大师Knuth的哦。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　考虑下第一次杀人的时候，编号为k = (m - 1) % n的同学挂了，那我们从k + 1重新从0开始编号
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　k + 1 ==&gt; 0
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　k + 2 ==&gt; 1
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　……
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　……
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　k - 2 ==&gt; n - 2
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　k - 1 ==&gt; n - 1
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　好了，剩余的n - 1个同学又组成了一个新的Joseph环，对新环来说，编号k = (m - 1) % (n - 1)的同学会挂，如此下去，这里面似乎有某种规律可寻。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　考虑到不会死的同学一直不会被杀(废话)，我们设i个同学时的不会挂的同学的编号(即解)为x，那么当死掉一个同学剩余i - 1个同学的时候，x仍然不会被杀，但此时的x已经由编号变换变成了x’，即x’是i - 1的情况时的解！一直推下去直到i - (i - 1)即1的情况，那1的时候解明显是0嘛！(注意编号是从0开始的)，倒推回来，那问题不就解决了么！
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　好了，分析清楚了剩下的就只是数学推导了，这个我比较烦，直接给公式吧：
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　向下变换：x’ = (x - (k + 1)) % i;
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　向上变换：x = (x + k + 1) % i;
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　其中：　　k = (m - 1) % i;
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　带入可得：x = (x’ + m) % i;
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　OK，现在应该很好写代码了：
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">　　
				<code style="line-height:25px"><span style="font-family:新宋体; line-height:25px">int Joseph(int n,int m)<br style="line-height:25px" />
				　　{<br style="line-height:25px" />
				&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px"><span>&nbsp;</span>　　int i,result;<br style="line-height:25px" />
				&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px"><br style="line-height:25px" />
				&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px"><span>&nbsp;</span>　　for (i = 1,result = 0;i &lt; n;i++)<br style="line-height:25px" />
				&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px"><span>&nbsp;</span>　　result = (result + m) % i;<br style="line-height:25px" />
				&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px">&nbsp;<wbr style="line-height:25px"><span>&nbsp;</span>　　return result;<br style="line-height:25px" />
				　　}</wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></wbr></span></code>
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　现在再来看看HOJ上的两道Joseph问题，有了上面的基础，只要做个脑筋急转弯就行了。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　先看第一道(1016)，说Joseph把杀人的秘密传出来之后，很多聪明的或者还不够聪明的程序员都能够选择合适的位置使自己存活下来，但 Joseph的表兄，一个恶毒的一直致力于将世界上所有愚蠢的程序员杀光光的人，又发明出一种新的玩法来杀人，如果你又很不幸的站在了这n个人里面，你会选择哪个位置站下。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　Joseph表兄的新玩法的特点是m是不再是给定的一个确定的数值，而且是按照素数从小到大变动的，即取值依次为{2, 3, 5, 7, ……}。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　当然你仍然选择不会挂的位置站下。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　题目已经说得很清楚了，不难想到我们只要在递推的过程中更新m的值就行了，m属于素数的集合，所以我们只要开个数组保存有用的素数，然后使m在其中取值就行了。
			</p>
			<p style="padding-bottom:0px; padding-top:0px; padding-left:0px; margin:0px 0px 10px; line-height:25px; padding-right:0px">
				　　还有一点值得注意的是由于第1次杀人的时候取的值是prime[0]，第2次是prime[1]，……第n次是prime[n-1]，所以往回推的时候第1次取的应该是prime[n-1]，第2次是prime[n-2]，……则第i次应该是prime[n-i]。
			</p>
		</div>
	</div>
</div>
<br />

<pre class="cpp" name="code">#include &lt;iostream&gt;

const int SIZE = 3502;
int prime[SIZE];
int result[SIZE];
int isprime(int n)
{
    for (int i=3 ;i*i&lt;=n;i+=2)
    {
        if (n%i==0)
            return 0;
    }
    return 1;
}
int main()
{
    int prime[SIZE];
    int result[SIZE];
    int i,j;

    for(i = 3,j = 1,prime[0] = 2;j &lt; SIZE;i += 2)
        if(isprime(i))
            prime[j++] = i;
    for(i = 1;i &lt; SIZE;i++)
        for(j = 1,result[i] = 0;j &lt;= i;j++)
            result[i] = (result[i] + prime[i-j]) % j;
    while(scanf(&quot;%d&quot;,&amp;i) == 1 &amp;&amp; i)
    {
        printf(&quot;%d\n&quot;,result[i] + 1);
    }
    return 0;
}
</pre>
<br />