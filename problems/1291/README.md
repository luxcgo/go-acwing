## [1291\. 轻拍牛头](https://www.acwing.com/problem/content/1293/)

### 题目

今天是贝茜的生日，为了庆祝自己的生日，贝茜邀你来玩一个游戏．

贝茜让 $N$ 头奶牛（编号 $1$ 到 $N$）坐成一个圈。

除了 $1$ 号与 $N$ 号奶牛外，$i$ 号奶牛与 $i-1$ 号和 $i+1$ 号奶牛相邻，$N$ 号奶牛与 $1$ 号奶牛相邻。

农夫约翰用很多纸条装满了一个桶，每一张纸条中包含一个 $1$ 到 $1000000$ 之间的数字。

接着每一头奶牛 $i$ 从桶中取出一张纸条，纸条上的数字用 $A\_i$ 表示。

所有奶牛都选取完毕后，每头奶牛轮流走上一圈，当走到一头奶牛身旁时，如果自己手中的数字能够被该奶牛手中的数字整除，则拍打该牛的头。

牛们希望你帮助他们确定，每一头奶牛需要拍打的牛的数量。

即共有 $N$ 个整数 $A\_1,A\_2,…,A\_N$，对于每一个数 $A\_i$，求其他的数中有多少个是它的约数。

### 输入格式

第一行包含整数 $N$。

接下来 $N$ 行，每行包含一个整数 $A\_i$。

### 输出格式

共 $N$ 行，第 $i$ 行的数字为第 $i$ 头牛需要拍打的牛的数量。

### 数据范围

$1 \\le N \\le 10^5$,

$1 \\le A\_i \\le 10^6$

### 输入样例：

```
5
2
1
2
3
4
```

### 输出样例：

```
2
0
2
1
3
```

### 题解
