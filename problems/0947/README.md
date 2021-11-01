## [947\. 文本编辑器](https://www.acwing.com/problem/content/949/)

### 题目

很久很久以前，DOS3.x 的程序员们开始对 EDLIN 感到厌倦。

于是，人们开始纷纷改用自己写的文本编辑器⋯⋯

多年之后，出于偶然的机会，小明找到了当时的一个编辑软件。

进行了一些简单的测试后，小明惊奇地发现：那个软件每秒能够进行上万次编辑操作（当然，你不能手工进行这样的测试）！

于是，小明废寝忘食地想做一个同样的东西出来。

你能帮助他吗？

为了明确目标，小明对“文本编辑器”做了一个抽象的定义：

文本：由 $0$ 个或多个 ASCII 码在闭区间 $[32, 126]$ 内的字符构成的序列。

光标：在一段文本中用于指示位置的标记，可以位于文本首部，文本尾部或文本的某两个字符之间。

文本编辑器：由一段文本和该文本中的一个光标组成的，支持如下操作的数据结构。如果这段文本为空，我们就说这个文本编辑器是空的。

![QQ截图20190902161827.png](https://cdn.acwing.com/media/article/image/2019/09/02/19_422b1060cd-QQ截图20190902161827.png)

比如一个空的文本编辑器依次执行操作 `INSERT(13, “Balanced tree”)， MOVE(2)，DELETE(5)，NEXT()，INSERT(7, “editor”)，MOVE(0)，GET(16)` 后，会输出 `Bad editor tree`。

你的任务是：

- 建立一个空的文本编辑器。
- 从输入文件中读入一些操作并执行。
- 对所有执行过的`GET` 操作，将指定的内容写入输出文件。

### 输入格式

输入文件的第一行是指令条数 $t$，以下是需要执行的 $t$ 个操作。

其中：

- 为了使输入文件便于阅读， `Insert` 操作的字符串中可能会插入一些回车符，请忽略掉它们（如果难以理解这句话，可以参照样例）。
- 除了回车符之外，输入文件的所有字符的 ASCII 码都在闭区间 $[32, 126]$ 内。且行尾没有空格。

这里我们有如下假定：

- `MOVE` 操作不超过 $50000$ 个， `INSERT` 和 `DELETE` 操作的总个数不超过 $4000$， `PREV` 和 `NEXT` 操作的总个数不超过 $200000$。
- 所有`INSERT` 插入的字符数之和不超过 $2M$（$1M=1024 \\times 1024$ 字节），正确的输出文件长度不超过 $3M$ 字节。
- `DELETE` 操作和 `GET` 操作执行时光标后必然有足够的字符。 `MOVE、PREV、NEXT` 操作必然不会试图把光标移动到非法位置。
- 输入文件没有错误。

### 输出格式

输出文件的每行依次对应输入文件中每条 `Get` 指令的输出。

### 输入样例：

```
15
Insert 26
abcdefghijklmnop
qrstuv wxy
Move 15
Delete 11
Move 5
Insert 1
^
Next
Insert 1
_
Next
Next
Insert 4
.\/.
Get 4
Prev
Insert 1
^
Move 0
Get 22
```

### 输出样例：

```
.\/.
abcde^_^f.\/.ghijklmno
```

### 题解
