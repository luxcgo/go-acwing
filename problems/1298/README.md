## [1298\. 曹冲养猪](https://www.acwing.com/problem/content/1300/)

### 题目

自从曹冲搞定了大象以后，曹操就开始琢磨让儿子干些事业，于是派他到中原养猪场养猪，可是曹冲很不高兴，于是在工作中马马虎虎，有一次曹操想知道母猪的数量，于是曹冲想狠狠耍曹操一把。

举个例子，假如有 $16$ 头母猪，如果建了 $3$ 个猪圈，剩下 $1$ 头猪就没有地方安家了；如果建造了 $5$ 个猪圈，但是仍然有 $1$ 头猪没有地方去；如果建造了 $7$ 个猪圈，还有 $2$ 头没有地方去。

你作为曹总的私人秘书理所当然要将准确的猪数报给曹总，你该怎么办？

### 输入格式

第一行包含一个整数 $n$，表示建立猪圈的次数；

接下来 $n$ 行，每行两个整数 $a\_i,b\_i$，表示建立了 $a\_i$ 个猪圈，有 $b\_i$ 头猪没有去处。

你可以假定 $a\_i,a\_j$ 互质。

### 输出格式

输出仅包含一个正整数，即为曹冲至少养猪的数目。

### 数据范围

$1 \\le n \\le 10$,

$1 \\le b\_i \\le a\_i \\le 1100000$

所有$a\_i$的乘积不超过 $10^{18}$

### 输入样例：

```
3
3 1
5 1
7 2
```

### 输出样例：

```
16
```

### 题解
