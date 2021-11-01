## [2476\. 树套树](https://www.acwing.com/problem/content/2478/)

### 题目

请你写出一种数据结构，来维护一个长度为 $n$ 的数列，其中需要提供以下操作：

1. `1 l r x`，查询整数 $x$ 在区间 $[l,r]$ 内的排名。
2. `2 l r k`，查询区间 $[l,r]$ 内排名为 $k$ 的值。
3. `3 pos x`，将 $pos$ 位置的数修改为 $x$。
4. `4 l r x`，查询整数 $x$ 在区间 $[l,r]$ 内的前驱(前驱定义为小于 $x$，且最大的数)。
5. `5 l r x`，查询整数 $x$ 在区间 $[l,r]$ 内的后继(后继定义为大于 $x$，且最小的数)。

数列中的位置从左到右依次标号为 $1 \\sim n$。

区间 $[l,r]$ 表示从位置 $l$ 到位置 $r$ 之间（包括两端点）的所有数字。

区间内排名为 $k$ 的值指区间内从小到大排在第 $k$ 位的数值。（位次从 $1$ 开始）

### 输入格式

第一行包含两个整数 $n,m$，表示数列长度以及操作次数。

第二行包含 $n$ 个整数，表示有序数列。

接下来 $m$ 行，每行包含一个操作指令，格式如题目所述。

### 输出格式

对于所有操作 $1,2,4,5$，每个操作输出一个查询结果，每个结果占一行。

### 数据范围

$1 \\le n,m \\le 5 \\times 10^4$,

$1 \\le l \\le r \\le n$,

$1 \\le pos \\le n$,

$1 \\le k \\le r-l+1$,

$0 \\le x \\le 10^8$,

有序数列中的数字始终满足在 $[0,10^8]$ 范围内，

数据保证所有操作一定合法，所有查询一定有解。

### 输入样例：

```
9 6
4 2 2 1 9 4 0 1 1
2 1 4 3
3 4 10
2 1 4 3
1 2 5 9
4 3 9 5
5 2 8 5
```

### 输出样例：

```
2
4
3
4
9
```

### 题解
