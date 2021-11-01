## [2192\. 运输问题](https://www.acwing.com/problem/content/2194/)

### 题目

$W$ 公司有 $m$ 个仓库和 $n$ 个零售商店。

第 $i$ 个仓库有 $a\_i$ 个单位的货物；第 $j$ 个零售商店需要 $b\_j$ 个单位的货物。

货物供需平衡，即$\\sum\\limits\_{i=1}^{m}a\_i=\\sum\\limits\_{j=1}^{n}b\_j$。

从第 $i$ 个仓库运送每单位货物到第 $j$ 个零售商店的费用为 $c\_{ij}$。

试设计一个将仓库中所有货物运送到零售商店的运输方案。

对于给定的 $m$ 个仓库和 $n$ 个零售商店间运送货物的费用，计算最优运输方案和最差运输方案。

### 输入格式

第 $1$ 行有 $2$ 个正整数 $m$ 和 $n$，分别表示仓库数和零售商店数。

接下来的一行中有 $m$ 个正整数 $a\_i$，表示第 $i$ 个仓库有 $a\_i$ 个单位的货物。

再接下来的一行中有 $n$ 个正整数 $b\_j$，表示第 $j$ 个零售商店需要 $b\_j$ 个单位的货物。

接下来的 $m$ 行，每行有 $n$ 个整数，表示从第 $i$ 个仓库运送每单位货物到第 $j$ 个零售商店的费用 $c\_{ij}$。

### 输出格式

第一行输出最少运输费用。

第二行输出最多运输费用。

### 数据范围

$1 \\le m \\le 100$,

$1 \\le n \\le 50$,

$1 \\le a\_i \\le 30000$,

$1 \\le b\_i \\le 60000$,

$1 \\le c\_{ij} \\le 1000$

### 输入样例：

```
2 3
220 280
170 120 210
77 39 105
150 186 122
```

### 输出样例：

```
48500
69140
```

### 题解
