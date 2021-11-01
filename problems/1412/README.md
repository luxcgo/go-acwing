## [1412\. 邮政货车](https://www.acwing.com/problem/content/1414/)

### 题目

奶牛们厌倦了美丽的田地，它们搬到了一个新的郊区居住。

郊区可以看作一个街道连接而成的矩形网格图，其中西北角是一个邮局。

整个郊区共有四个东西延伸的街道和 $N$ 个南北延伸的街道。

例如，下图是一个 $N = 5$ 时的郊区道路图，其中左上角的黑点代表邮局。

![postal1.gif](https://cdn.acwing.com/media/article/image/2020/03/05/19_197c405a5e-postal1.gif)

每天，邮政货车都会从邮局出发，在郊区行驶，然后返回邮局，在这一过程中，每个十字路口（包括边界或拐角处）都需要恰好经过一次。

邮政公司的管理人员想知道可以为邮政货车建立多少条不同的路线（当然，在此计数中，路线方向很重要）。

例如，下图显示了上述郊区的两条满足条件的路线：

![postal2.gif](https://cdn.acwing.com/media/article/image/2020/03/05/19_ac50ce3c5e-postal2.gif)

再例如，下图给出了当 $N = 3$ 时，所有四种可能的行进路线：

![postal3.gif](https://cdn.acwing.com/media/article/image/2020/03/05/19_ff70619a5e-postal3.gif)

请你编写一个程序，在给定街道数量的情况下，计算共有多少条不同的行进路线。

### 输入格式

共一行，包含一个整数 $N$。

### 输出格式

输出一个整数，表示总路线条数。

### 数据范围

$1 \\le N \\le 1000$

### 输入样例：

```
4
```

### 输出样例：

```
12
```

### 题解
