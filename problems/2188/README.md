## [2188\. 无源汇上下界可行流](https://www.acwing.com/problem/content/2190/)

### 题目

给定一个包含 $n$ 个点 $m$ 条边的有向图，每条边都有一个流量下界和流量上界。

求一种可行方案使得在所有点满足流量平衡条件的前提下，所有边满足流量限制。

### 输入格式

第一行包含两个整数 $n$ 和 $m$。

接下来 $m$ 行，每行包含四个整数 $a,b,c,d$ 表示点 $a$ 和 $b$ 之间存在一条有向边，该边的流量下界为 $c$，流量上界为 $d$。

点编号从 $1$ 到 $n$。

### 输出格式

如果存在可行方案，则第一行输出 `YES`，接下来 $m$ 行，每行输出一个整数，其中第 $i$ 行的整数表示输入的第 $i$ 条边的流量。

如果不存在可行方案，直接输出一行 `NO`。

如果可行方案不唯一，则输出任意一种方案即可。

### 数据范围

$1 \\le n \\le 200$,

$1 \\le m \\le 10200$,

$1 \\le a,b \\le n$,

$0 \\le c \\le d \\le 10000$

### 输入样例1：

```
4 6
1 2 1 3
2 3 1 3
3 4 1 3
4 1 1 3
1 3 1 3
4 2 1 3
```

### 输出样例1：

```
YES
1
2
3
2
1
1
```

### 输入样例2：

```
4 6
1 2 1 2
2 3 1 2
3 4 1 2
4 1 1 2
1 3 1 2
4 2 1 2
```

### 输出样例2：

```
NO
```

### 题解
