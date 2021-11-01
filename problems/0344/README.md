## [344\. 观光之旅](https://www.acwing.com/problem/content/346/)

### 题目

给定一张无向图，求图中一个至少包含 $3$ 个点的环，环上的节点不重复，并且环上的边的长度之和最小。

该问题称为无向图的最小环问题。

你需要输出最小环的方案，若最小环不唯一，输出任意一个均可。

### 输入格式

第一行包含两个整数 $N$ 和 $M$，表示无向图有 $N$ 个点，$M$ 条边。

接下来 $M$ 行，每行包含三个整数 $u，v，l$，表示点 $u$ 和点 $v$ 之间有一条边，边长为 $l$。

### 输出格式

输出占一行，包含最小环的所有节点（按顺序输出），如果不存在则输出 `No solution.`。

### 数据范围

$1 \\le N \\le 100$,

$1 \\le M \\le 10000$,

$1 \\le l < 500$

### 输入样例：

```
5 7
1 4 1
1 3 300
3 1 10
1 2 16
2 3 100
2 5 15
5 3 20
```

### 输出样例：

```
1 3 5 2
```

### 题解
