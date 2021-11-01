## [1015\. 摘花生](https://www.acwing.com/problem/content/1017/)

### 题目

Hello Kitty想摘点花生送给她喜欢的米老鼠。

她来到一片有网格状道路的矩形花生地(如下图)，从西北角进去，东南角出来。

地里每个道路的交叉点上都有种着一株花生苗，上面有若干颗花生，经过一株花生苗就能摘走该它上面所有的花生。

Hello Kitty只能向东或向南走，不能向西或向北走。

问Hello Kitty最多能够摘到多少颗花生。

![1.gif](https://cdn.acwing.com/media/article/image/2019/09/12/19_a8509f26d5-1.gif)

### 输入格式

第一行是一个整数T，代表一共有多少组数据。

接下来是T组数据。

每组数据的第一行是两个整数，分别代表花生苗的行数R和列数 C。

每组数据的接下来R行数据，从北向南依次描述每行花生苗的情况。每行数据有C个整数，按从西向东的顺序描述了该行每株花生苗上的花生数目M。

### 输出格式

对每组输入数据，输出一行，内容为Hello Kitty能摘到得最多的花生颗数。

### 数据范围

$1 \\le T \\le 100$,

$1 \\le R,C \\le 100$,

$0 \\le M \\le 1000$

### 输入样例：

```
2
2 2
1 1
3 4
2 3
2 3 4
1 6 5
```

### 输出样例：

```
8
16
```

### 题解
