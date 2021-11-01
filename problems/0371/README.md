## [371\. 牧师约翰最忙碌的一天](https://www.acwing.com/problem/content/373/)

### 题目

牧师约翰在 $9$ 月 $1$ 日这天非常的忙碌。

有 $N$ 对情侣在这天准备结婚，每对情侣都预先计划好了婚礼举办的时间，其中第 $i$ 对情侣的婚礼从时刻 $S\_i$ 开始，到时刻 $T\_i$ 结束。

婚礼有一个必须的仪式：站在牧师面前聆听上帝的祝福。

这个仪式要么在婚礼开始时举行，要么在结束时举行。

第 $i$ 对情侣需要 $D\_i$ 分钟完成这个仪式，即必须选择 $S\_i \\sim S\_i+D\_i$ 或 $T\_i-D\_i \\sim T\_i$ 两个时间段之一。

牧师想知道他能否满足每场婚礼的要求，即给每对情侣安排$S\_i \\sim S\_i+D\_i$ 或 $T\_i-D\_i \\sim T\_i$，使得这些仪式的时间段不重叠。

若能满足，还需要帮牧师求出任意一种具体方案。

注意，约翰不能同时主持两场婚礼，且 **所有婚礼的仪式均发生在 $9$ 月 $1$ 日当天**。

如果一场仪式的结束时间与另一场仪式的开始时间相同，则不算重叠。

例如：一场仪式安排在 $08:00 \\sim 09:00$，另一场仪式安排在 $09:00 \\sim 10:00$，则不认为两场仪式出现重叠。

### 输入格式

第一行包含整数 $N$。

接下来 $N$ 行，每行包含 $S\_i,T\_i,D\_i$，其中 $S\_i$ 和 $T\_i$ 是 $hh:mm$ 形式。

### 输出格式

第一行输出能否满足，能则输出 `YES`，否则输出 `NO`。

接下来 $N$ 行，每行给出一个具体时间段安排。

### 数据范围

$1 \\le N \\le 1000$

### 输入样例：

```
2
08:00 09:00 30
08:15 09:00 20
```

### 输出样例：

```
YES
08:00 08:30
08:40 09:00
```

### 题解
