## [243\. 一个简单的整数问题2](https://www.acwing.com/problem/content/244/)

### 题目

给定一个长度为 $N$ 的数列 $A$，以及 $M$ 条指令，每条指令可能是以下两种之一：

1. `C l r d`，表示把 $A[l],A[l+1],…,A[r]$ 都加上 $d$。
2. `Q l r`，表示询问数列中第 $l \\sim r$ 个数的和。

对于每个询问，输出一个整数表示答案。

### 输入格式

第一行两个整数 $N,M$。

第二行 $N$ 个整数 $A[i]$。

接下来 $M$ 行表示 $M$ 条指令，每条指令的格式如题目描述所示。

### 输出格式

对于每个询问，输出一个整数表示答案。

每个答案占一行。

### 数据范围

$1 \\le N,M \\le 10^5$,

$\|d\| \\le 10000$,

$\|A[i]\| \\le 10^9$

### 输入样例：

```
10 5
1 2 3 4 5 6 7 8 9 10
Q 4 4
Q 1 10
Q 2 4
C 3 6 3
Q 2 4
```

### 输出样例：

```
4
55
9
15
```

### 题解
