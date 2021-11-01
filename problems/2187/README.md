## [2187\. 星际转移问题](https://www.acwing.com/problem/content/2189/)

### 题目

由于人类对自然资源的消耗，人们意识到大约在 $2300$ 年之后，地球就不能再居住了。

于是在月球上建立了新的绿地，以便在需要时移民。

令人意想不到的是，$2177$ 年冬由于未知的原因，地球环境发生了连锁崩溃，人类必须在最短的时间内迁往月球。

现有 $n$ 个太空站（编号 $1 \\sim n$）位于地球与月球之间，且有 $m$ 艘公共交通太空船在其间来回穿梭。

每个太空站可容纳无限多的人，而每艘太空船 $i$ 只可容纳 $H[i]$ 个人。

每艘太空船将周期性地停靠一系列的太空站，例如：$(1，3，4)$ 表示该太空船将周期性地停靠太空站 $134134134…$。

每一艘太空船从一个太空站驶往任一太空站耗时均为 $1$。

人们只能在太空船停靠太空站(或月球、地球)时上、下船。

初始时所有人全在地球上，太空船全在初始站，即行驶周期中的第一个站。

试设计一个算法，找出让所有人尽快地全部转移到月球上的运输方案。

### 输入格式

第 $1$ 行有 $3$ 个正整数 $n$（太空站个数），$m$（太空船个数）和 $k$（需要运送的地球上的人的个数）。

接下来的 $m$ 行给出太空船的信息。第 $i+1$ 行说明太空船 $p\_i$。第 $1$ 个数表示 $p\_i$ 可容纳的人数 $H[p\_i]$；第 $2$ 个数表示 $p\_i$ 一个周期停靠的太空站个数 $r$；随后 $r$ 个数是停靠的太空站的编号 $(S\_{i1},S\_{i2},…,S\_{ir})$，地球用 $0$ 表示，月球用 $-1$ 表示。

时刻 $0$ 时，所有太空船都在初始站，然后开始运行。

在时刻 $1，2，3…$ 等正点时刻各艘太空船停靠相应的太空站。

人只有在 $0,1,2…$ 等正点时刻才能上下太空船。

### 输出格式

输出让所有人尽快地全部转移到月球上的最短用时。

如果无解，则输出 $0$。

### 数据范围

$1 \\le n \\le 13$,

$1 \\le m \\le 20$,

$1 \\le k \\le 50$,

$1 \\le r \\le n+2$,

### 输入样例：

```
2 2 1
1 3 0 1 2
1 3 1 2 -1
```

### 输出样例：

```
5
```

### 题解
