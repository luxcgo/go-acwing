# go-acwing

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/luxcgo/go-acwing?tab=doc)
[![GoReport](https://goreportcard.com/badge/github.com/luxcgo/go-acwing?style=for-the-badge)](https://goreportcard.com/report/github.com/luxcgo/go-acwing)
[![Sourcegraph](https://img.shields.io/badge/view%20on-Sourcegraph-brightgreen.svg?style=for-the-badge&logo=sourcegraph)](https://sourcegraph.com/github.com/go-ini/ini)

**acwing.com solutions in golang** 43/459

> 推荐安装 [MathJax Plugin for Github](https://chrome.google.com/webstore/detail/mathjax-plugin-for-github/ioemnmodlmafdkllaclgeombjnmnbima) 以获得最佳的观看体验

| # | 标题 | 题解 | 难度 | 前置题目 | 前置知识 | 本题知识 |
| - | - | - | - | - | - | - |
| 0002 | 01背包问题 | [Go](problems/0002) | 基础 | 0000 | 无 | 动态规划-背包问题 |
| 0003 | 完全背包问题 | [Go](problems/0003) | 基础 | 0002 | 01背包 | 动态规划-背包问题 |
| 0004 | 多重背包问题I | [Go](problems/0004) | 基础 | 0003 | 完全背包 | 动态规划-背包问题 |
| 0005 | 多重背包问题II | [Go](problems/0005) | 基础 | 0004 | 01背包 | 动态规划-背包问题 |
| 0009 | 分组背包问题 | [Go](problems/0009) | 基础 | 0005 | 01背包 | 动态规划-背包问题 |
| 0091 | 最短Hamilton路径 | [Go](problems/0091) | 基础 | 0291 | 位运算 | 动态规划-状态压缩DP |
| 0104 | 货仓选址 | [Go](problems/0104) | 基础 |  |  |  |
| 0125 | 耍杂技的牛 | [Go](problems/0125) | 基础 |  |  |  |
| 0143 | 最大异或对 | [Go](problems/0143) | 基础 |  |  |  |
| 0148 | 合并果子 | [Go](problems/0148) | 基础 | 0907 | 小根堆 | 贪心-Huffman树 |
| 0154 | 滑动窗口 | [Go](problems/0154) | 基础 |  |  |  |
| 0204 | 表达整数的奇怪方式 | [Go](problems/0204) | 基础 |  |  |  |
| 0240 | 食物链 | [Go](problems/0240) | 基础 |  |  |  |
| 0282 | 石子合并 | [Go](problems/0282) | 基础 | 0899 | 前缀和 | 动态规划-区间DP |
| 0285 | 没有上司的舞会 | [Go](problems/0285) | 基础 |  |  |  |
| 0291 | 蒙德里安的梦想 | [Go](problems/0291) | 基础 | 0338 | 简单数论 | 动态规划-状态压缩DP |
| 0338 | 计数问题 | [Go](problems/0338) | 基础 |  |  |  |
| 0785 | 快速排序 | [Go](problems/0785) | 基础 | 0000 | 语法 | 基础算法-快速排序 |
| 0786 | 第k个数 | [Go](problems/0786) | 基础 | 0785 | 快排 | 基础算法-快速排序 |
| 0787 | 归并排序 | [Go](problems/0787) | 基础 | 0786 | 分治思想 | 基础算法-归并排序 |
| 0788 | 逆序对的数量 | [Go](problems/0788) | 基础 | 0787 | 归并排序 | 基础算法-归并排序 |
| 0789 | 数的范围 | [Go](problems/0789) | 基础 | 0788 | 语法 | 基础算法-二分 |
| 0790 | 数的三次方根 | [Go](problems/0790) | 基础 |  |  |  |
| 0791 | 高精度加法 | [Go](problems/0791) | 基础 |  |  |  |
| 0792 | 高精度减法 | [Go](problems/0792) | 基础 |  |  |  |
| 0793 | 高精度乘法 | [Go](problems/0793) | 基础 |  |  |  |
| 0794 | 高精度除法 | [Go](problems/0794) | 基础 |  |  |  |
| 0795 | 前缀和 | [Go](problems/0795) | 基础 | 0794 | 语法 | 基础算法-前缀和 |
| 0796 | 子矩阵的和 | [Go](problems/0796) | 基础 | 0795 | 前缀和 | 基础算法-前缀和 |
| 0797 | 差分 | [Go](problems/0797) | 基础 | 0796 | 前缀和 | 基础算法-差分 |
| 0798 | 差分矩阵 | [Go](problems/0798) | 基础 | 0797 | 一维差分 | 基础算法-差分 |
| 0799 | 最长连续不重复子序列 | [Go](problems/0799) | 基础 | 0798 | 语法 | 基础算法-双指针 |
| 0800 | 数组元素的目标和 | [Go](problems/0800) | 基础 | 0799 | 双指针 | 基础算法-双指针 |
| 0801 | 二进制中1的个数 | [Go](problems/0801) | 基础 |  |  |  |
| 0802 | 区间和 | [Go](problems/0802) | 基础 | 0801 | 二分、前缀和 | 基础算法-离散化 |
| 0803 | 区间合并 | [Go](problems/0803) | 基础 | 0802 | 语法 | 基础算法-区间合并 |
| 0826 | 单链表 | [Go](problems/0826) | 基础 | 0000 | 语法 | 数据结构-单链表 |
| 0827 | 双链表 | [Go](problems/0827) | 基础 | 0826 | 单链表 | 数据结构-双链表 |
| 0828 | 模拟栈 | [Go](problems/0828) | 基础 | 0827 | 数组 | 数据结构-栈 |
| 0829 | 模拟队列 | [Go](problems/0829) | 基础 | 0828 | 数组 | 数据结构-队列 |
| 0830 | 单调栈 | [Go](problems/0830) | 基础 | 0829 | 栈 | 数据结构-单调栈 |
| 0831 | KMP字符串 | [Go](problems/0831) | 基础 |  |  |  |
| 0835 | Trie字符串统计 | [Go](problems/0835) | 基础 | 0831 | 邻接表 | 数据结构-Trie |
| 0836 | 合并集合 | [Go](problems/0836) | 基础 | 0143 | 树 | 数据结构-并查集 |
| 0837 | 连通块中点的数量 | [Go](problems/0837) | 基础 |  |  |  |
| 0838 | 堆排序 | [Go](problems/0838) | 基础 | 0240 | 语法 | 数据结构-堆 |
| 0839 | 模拟堆 | [Go](problems/0839) | 基础 | 0838 | 堆排序 | 数据结构-堆 |
| 0840 | 模拟散列表 | [Go](problems/0840) | 基础 | 0839 | 单链表 | 数据结构-哈希表 |
| 0841 | 字符串哈希 | [Go](problems/0841) | 基础 |  |  |  |
| 0842 | 排列数字 | [Go](problems/0842) | 基础 | 0841 | 递归 | 搜索与图论-DFS |
| 0843 | n-皇后问题 | [Go](problems/0843-n) | 基础 | 0842 | 数学 | 搜索与图论-DFS |
| 0844 | 走迷宫 | [Go](problems/0844) | 基础 | 0843 | 队列 | 搜索与图论-BFS |
| 0845 | 八数码 | [Go](problems/0845) | 基础 |  |  |  |
| 0846 | 树的重心 | [Go](problems/0846) | 基础 |  |  |  |
| 0847 | 图中点的层次 | [Go](problems/0847) | 基础 |  |  |  |
| 0848 | 有向图的拓扑序列 | [Go](problems/0848) | 基础 |  |  |  |
| 0849 | Dijkstra求最短路I | [Go](problems/0849) | 基础 |  |  |  |
| 0850 | Dijkstra求最短路II | [Go](problems/0850) | 基础 |  |  |  |
| 0851 | spfa求最短路 | [Go](problems/0851) | 基础 |  |  |  |
| 0852 | spfa判断负环 | [Go](problems/0852) | 基础 |  |  |  |
| 0853 | 有边数限制的最短路 | [Go](problems/0853) | 基础 |  |  |  |
| 0854 | Floyd求最短路 | [Go](problems/0854) | 基础 |  |  |  |
| 0858 | Prim算法求最小生成树 | [Go](problems/0858) | 基础 |  |  |  |
| 0859 | Kruskal算法求最小生成树 | [Go](problems/0859) | 基础 |  |  |  |
| 0860 | 染色法判定二分图 | [Go](problems/0860) | 基础 |  |  |  |
| 0861 | 二分图的最大匹配 | [Go](problems/0861) | 基础 |  |  |  |
| 0866 | 试除法判定质数 | [Go](problems/0866) | 基础 |  |  |  |
| 0867 | 分解质因数 | [Go](problems/0867) | 基础 |  |  |  |
| 0868 | 筛质数 | [Go](problems/0868) | 基础 |  |  |  |
| 0869 | 试除法求约数 | [Go](problems/0869) | 基础 |  |  |  |
| 0870 | 约数个数 | [Go](problems/0870) | 基础 |  |  |  |
| 0871 | 约数之和 | [Go](problems/0871) | 基础 |  |  |  |
| 0872 | 最大公约数 | [Go](problems/0872) | 基础 |  |  |  |
| 0873 | 欧拉函数 | [Go](problems/0873) | 基础 |  |  |  |
| 0874 | 筛法求欧拉函数 | [Go](problems/0874) | 基础 |  |  |  |
| 0875 | 快速幂 | [Go](problems/0875) | 基础 |  |  |  |
| 0876 | 快速幂求逆元 | [Go](problems/0876) | 基础 |  |  |  |
| 0877 | 扩展欧几里得算法 | [Go](problems/0877) | 基础 |  |  |  |
| 0878 | 线性同余方程 | [Go](problems/0878) | 基础 |  |  |  |
| 0883 | 高斯消元解线性方程组 | [Go](problems/0883) | 基础 |  |  |  |
| 0884 | 高斯消元解异或线性方程组 | [Go](problems/0884) | 基础 |  |  |  |
| 0885 | 求组合数I | [Go](problems/0885) | 基础 |  |  |  |
| 0886 | 求组合数II | [Go](problems/0886) | 基础 |  |  |  |
| 0887 | 求组合数III | [Go](problems/0887) | 基础 |  |  |  |
| 0888 | 求组合数IV | [Go](problems/0888) | 基础 |  |  |  |
| 0889 | 满足条件的01序列 | [Go](problems/0889) | 基础 |  |  |  |
| 0890 | 能被整除的数 | [Go](problems/0890) | 基础 |  |  |  |
| 0891 | Nim游戏 | [Go](problems/0891) | 基础 |  |  |  |
| 0892 | 台阶-Nim游戏 | [Go](problems/0892-台阶) | 基础 |  |  |  |
| 0893 | 集合-Nim游戏 | [Go](problems/0893-集合) | 基础 |  |  |  |
| 0894 | 拆分-Nim游戏 | [Go](problems/0894-拆分) | 基础 |  |  |  |
| 0895 | 最长上升子序列 | [Go](problems/0895) | 基础 | 0898 | 无 | 动态规划-线性DP |
| 0896 | 最长上升子序列II | [Go](problems/0896) | 基础 |  |  |  |
| 0897 | 最长公共子序列 | [Go](problems/0897) | 基础 | 0896 | 无 | 动态规划-线性DP |
| 0898 | 数字三角形 | [Go](problems/0898) | 基础 | 0000 | 无 | 动态规划-线性DP |
| 0899 | 编辑距离 | [Go](problems/0899) | 基础 |  |  |  |
| 0900 | 整数划分 | [Go](problems/0900) | 基础 |  |  |  |
| 0901 | 滑雪 | [Go](problems/0901) | 基础 | 0285 | 递归 | 动态规划-记忆化搜索 |
| 0902 | 最短编辑距离 | [Go](problems/0902) | 基础 |  |  |  |
| 0905 | 区间选点 | [Go](problems/0905) | 基础 | 0000 | 无 | 贪心-区间问题 |
| 0906 | 区间分组 | [Go](problems/0906) | 基础 | 0908 | 小根堆 | 贪心-区间问题 |
| 0907 | 区间覆盖 | [Go](problems/0907) | 基础 | 0906 | 双指针 | 贪心-区间问题 |
| 0908 | 最大不相交区间数量 | [Go](problems/0908) | 基础 | 0905 | 无 | 贪心-区间问题 |
| 0913 | 排队打水 | [Go](problems/0913) | 基础 |  |  |  |
| 2816 | 判断子序列 | [Go](problems/2816) | 基础 |  |  |  |
| 3302 | 表达式求值 | [Go](problems/3302) | 基础 |  |  |  |
| 0006 | 多重背包问题III | [Go](problems/0006) | 提高 |  |  |  |
| 0007 | 混合背包问题 | [Go](problems/0007) | 提高 |  |  |  |
| 0008 | 二维费用的背包问题 | [Go](problems/0008) | 提高 |  |  |  |
| 0010 | 有依赖的背包问题 | [Go](problems/0010) | 提高 |  |  |  |
| 0011 | 背包问题求方案数 | [Go](problems/0011) | 提高 |  |  |  |
| 0012 | 背包问题求具体方案 | [Go](problems/0012) | 提高 |  |  |  |
| 0090 | 64位整数乘法 | [Go](problems/0090) | 提高 |  |  |  |
| 0095 | 费解的开关 | [Go](problems/0095) | 提高 |  |  |  |
| 0097 | 约数之和 | [Go](problems/0097) | 提高 |  |  |  |
| 0098 | 分形之城 | [Go](problems/0098) | 提高 |  |  |  |
| 0099 | 激光炸弹 | [Go](problems/0099) | 提高 |  |  |  |
| 0100 | 增减序列 | [Go](problems/0100) | 提高 |  |  |  |
| 0102 | 最佳牛围栏 | [Go](problems/0102) | 提高 |  |  |  |
| 0105 | 七夕祭 | [Go](problems/0105) | 提高 |  |  |  |
| 0106 | 动态中位数 | [Go](problems/0106) | 提高 |  |  |  |
| 0107 | 超快速排序 | [Go](problems/0107) | 提高 |  |  |  |
| 0113 | 特殊排序 | [Go](problems/0113) | 提高 |  |  |  |
| 0135 | 最大子序和 | [Go](problems/0135) | 提高 |  |  |  |
| 0164 | 可达性统计 | [Go](problems/0164) | 提高 |  |  |  |
| 0165 | 小猫爬山 | [Go](problems/0165) | 提高 |  |  |  |
| 0166 | 数独 | [Go](problems/0166) | 提高 |  |  |  |
| 0167 | 木棒 | [Go](problems/0167) | 提高 |  |  |  |
| 0168 | 生日蛋糕 | [Go](problems/0168) | 提高 |  |  |  |
| 0170 | 加成序列 | [Go](problems/0170) | 提高 |  |  |  |
| 0171 | 送礼物 | [Go](problems/0171) | 提高 |  |  |  |
| 0173 | 矩阵距离 | [Go](problems/0173) | 提高 |  |  |  |
| 0175 | 电路维修 | [Go](problems/0175) | 提高 |  |  |  |
| 0178 | 第K短路 | [Go](problems/0178) | 提高 |  |  |  |
| 0179 | 八数码 | [Go](problems/0179) | 提高 |  |  |  |
| 0180 | 排书 | [Go](problems/0180) | 提高 |  |  |  |
| 0181 | 回转游戏 | [Go](problems/0181) | 提高 |  |  |  |
| 0187 | 导弹防御系统 | [Go](problems/0187) | 提高 |  |  |  |
| 0188 | 武士风度的牛 | [Go](problems/0188) | 提高 |  |  |  |
| 0190 | 字串变换 | [Go](problems/0190) | 提高 |  |  |  |
| 0196 | 质数距离 | [Go](problems/0196) | 提高 |  |  |  |
| 0197 | 阶乘分解 | [Go](problems/0197) | 提高 |  |  |  |
| 0198 | 反素数 | [Go](problems/0198) | 提高 |  |  |  |
| 0200 | Hankson的趣味题 | [Go](problems/0200) | 提高 |  |  |  |
| 0201 | 可见的点 | [Go](problems/0201) | 提高 |  |  |  |
| 0202 | 最幸运的数字 | [Go](problems/0202) | 提高 |  |  |  |
| 0203 | 同余方程 | [Go](problems/0203) | 提高 |  |  |  |
| 0207 | 球形空间产生器 | [Go](problems/0207) | 提高 |  |  |  |
| 0208 | 开关问题 | [Go](problems/0208) | 提高 |  |  |  |
| 0214 | Devu和鲜花 | [Go](problems/0214) | 提高 |  |  |  |
| 0215 | 破译密码 | [Go](problems/0215) | 提高 |  |  |  |
| 0217 | 绿豆蛙的归宿 | [Go](problems/0217) | 提高 |  |  |  |
| 0218 | 扑克牌 | [Go](problems/0218) | 提高 |  |  |  |
| 0220 | 最大公约数 | [Go](problems/0220) | 提高 |  |  |  |
| 0222 | 青蛙的约会 | [Go](problems/0222) | 提高 |  |  |  |
| 0237 | 程序自动分析 | [Go](problems/0237) | 提高 |  |  |  |
| 0238 | 银河英雄传说 | [Go](problems/0238) | 提高 |  |  |  |
| 0239 | 奇偶游戏 | [Go](problems/0239) | 提高 |  |  |  |
| 0241 | 楼兰图腾 | [Go](problems/0241) | 提高 |  |  |  |
| 0242 | 一个简单的整数问题 | [Go](problems/0242) | 提高 |  |  |  |
| 0243 | 一个简单的整数问题2 | [Go](problems/0243) | 提高 |  |  |  |
| 0243 | 一个简单的整数问题2 | [Go](problems/0243) | 提高 |  |  |  |
| 0244 | 谜一样的牛 | [Go](problems/0244) | 提高 |  |  |  |
| 0245 | 你能回答这些问题吗 | [Go](problems/0245) | 提高 |  |  |  |
| 0246 | 区间最大公约数 | [Go](problems/0246) | 提高 |  |  |  |
| 0247 | 亚特兰蒂斯 | [Go](problems/0247) | 提高 |  |  |  |
| 0253 | 普通平衡树 | [Go](problems/0253) | 提高 |  |  |  |
| 0255 | 第K小数 | [Go](problems/0255) | 提高 |  |  |  |
| 0256 | 最大异或和 | [Go](problems/0256) | 提高 |  |  |  |
| 0257 | 关押罪犯 | [Go](problems/0257) | 提高 |  |  |  |
| 0265 | 营业额统计 | [Go](problems/0265) | 提高 |  |  |  |
| 0272 | 最长公共上升子序列 | [Go](problems/0272) | 提高 |  |  |  |
| 0275 | 传纸条 | [Go](problems/0275) | 提高 |  |  |  |
| 0278 | 数字组合 | [Go](problems/0278) | 提高 |  |  |  |
| 0292 | 炮兵阵地 | [Go](problems/0292) | 提高 |  |  |  |
| 0300 | 任务安排1 | [Go](problems/0300) | 提高 |  |  |  |
| 0301 | 任务安排2 | [Go](problems/0301) | 提高 |  |  |  |
| 0302 | 任务安排3 | [Go](problems/0302) | 提高 |  |  |  |
| 0303 | 运输小猫 | [Go](problems/0303) | 提高 |  |  |  |
| 0320 | 能量项链 | [Go](problems/0320) | 提高 |  |  |  |
| 0321 | 棋盘分割 | [Go](problems/0321) | 提高 |  |  |  |
| 0323 | 战略游戏 | [Go](problems/0323) | 提高 |  |  |  |
| 0327 | 玉米田 | [Go](problems/0327) | 提高 |  |  |  |
| 0340 | 通信线路 | [Go](problems/0340) | 提高 |  |  |  |
| 0341 | 最优贸易 | [Go](problems/0341) | 提高 |  |  |  |
| 0342 | 道路与航线 | [Go](problems/0342) | 提高 |  |  |  |
| 0343 | 排序 | [Go](problems/0343) | 提高 |  |  |  |
| 0344 | 观光之旅 | [Go](problems/0344) | 提高 |  |  |  |
| 0345 | 牛站 | [Go](problems/0345) | 提高 |  |  |  |
| 0346 | 走廊泼水节 | [Go](problems/0346) | 提高 |  |  |  |
| 0352 | 闇の連鎖 | [Go](problems/0352) | 提高 |  |  |  |
| 0356 | 次小生成树 | [Go](problems/0356) | 提高 |  |  |  |
| 0361 | 观光奶牛 | [Go](problems/0361) | 提高 |  |  |  |
| 0362 | 区间 | [Go](problems/0362) | 提高 |  |  |  |
| 0367 | 学校网络 | [Go](problems/0367) | 提高 |  |  |  |
| 0368 | 银河 | [Go](problems/0368) | 提高 |  |  |  |
| 0372 | 棋盘覆盖 | [Go](problems/0372) | 提高 |  |  |  |
| 0376 | 机器任务 | [Go](problems/0376) | 提高 |  |  |  |
| 0378 | 骑士放置 | [Go](problems/0378) | 提高 |  |  |  |
| 0379 | 捉迷藏 | [Go](problems/0379) | 提高 |  |  |  |
| 0383 | 观光 | [Go](problems/0383) | 提高 |  |  |  |
| 0393 | 雇佣收银员 | [Go](problems/0393) | 提高 |  |  |  |
| 0395 | 冗余路径 | [Go](problems/0395) | 提高 |  |  |  |
| 0396 | 矿场搭建 | [Go](problems/0396) | 提高 |  |  |  |
| 0423 | 采药 | [Go](problems/0423) | 提高 |  |  |  |
| 0426 | 开心的金明 | [Go](problems/0426) | 提高 |  |  |  |
| 0456 | 车站分级 | [Go](problems/0456) | 提高 |  |  |  |
| 0479 | 加分二叉树 | [Go](problems/0479) | 提高 |  |  |  |
| 0482 | 合唱队形 | [Go](problems/0482) | 提高 |  |  |  |
| 0487 | 金明的预算方案 | [Go](problems/0487) | 提高 |  |  |  |
| 0524 | 愤怒的小鸟 | [Go](problems/0524) | 提高 |  |  |  |
| 0529 | 宝藏 | [Go](problems/0529) | 提高 |  |  |  |
| 0532 | 货币系统 | [Go](problems/0532) | 提高 |  |  |  |
| 0734 | 能量石 | [Go](problems/0734) | 提高 |  |  |  |
| 0903 | 昂贵的聘礼 | [Go](problems/0903) | 提高 |  |  |  |
| 0904 | 虫洞 | [Go](problems/0904) | 提高 |  |  |  |
| 0920 | 最优乘车 | [Go](problems/0920) | 提高 |  |  |  |
| 1010 | 拦截导弹 | [Go](problems/1010) | 提高 |  |  |  |
| 1012 | 友好城市 | [Go](problems/1012) | 提高 |  |  |  |
| 1013 | 机器分配 | [Go](problems/1013) | 提高 |  |  |  |
| 1014 | 登山 | [Go](problems/1014) | 提高 |  |  |  |
| 1015 | 摘花生 | [Go](problems/1015) | 提高 |  |  |  |
| 1016 | 最大上升子序列和 | [Go](problems/1016) | 提高 |  |  |  |
| 1017 | 怪盗基德的滑翔翼 | [Go](problems/1017) | 提高 |  |  |  |
| 1018 | 最低通行费 | [Go](problems/1018) | 提高 |  |  |  |
| 1019 | 庆功会 | [Go](problems/1019) | 提高 |  |  |  |
| 1020 | 潜水员 | [Go](problems/1020) | 提高 |  |  |  |
| 1021 | 货币系统 | [Go](problems/1021) | 提高 |  |  |  |
| 1022 | 宠物小精灵之收服 | [Go](problems/1022) | 提高 |  |  |  |
| 1023 | 买书 | [Go](problems/1023) | 提高 |  |  |  |
| 1024 | 装箱问题 | [Go](problems/1024) | 提高 |  |  |  |
| 1027 | 方格取数 | [Go](problems/1027) | 提高 |  |  |  |
| 1049 | 大盗阿福 | [Go](problems/1049) | 提高 |  |  |  |
| 1052 | 设计密码 | [Go](problems/1052) | 提高 |  |  |  |
| 1053 | 修复DNA | [Go](problems/1053) | 提高 |  |  |  |
| 1057 | 股票买卖IV | [Go](problems/1057) | 提高 |  |  |  |
| 1058 | 股票买卖V | [Go](problems/1058) | 提高 |  |  |  |
| 1064 | 小国王 | [Go](problems/1064) | 提高 |  |  |  |
| 1068 | 环形石子合并 | [Go](problems/1068) | 提高 |  |  |  |
| 1069 | 凸多边形的划分 | [Go](problems/1069) | 提高 |  |  |  |
| 1072 | 树的最长路径 | [Go](problems/1072) | 提高 |  |  |  |
| 1073 | 树的中心 | [Go](problems/1073) | 提高 |  |  |  |
| 1074 | 二叉苹果树 | [Go](problems/1074) | 提高 |  |  |  |
| 1075 | 数字转换 | [Go](problems/1075) | 提高 |  |  |  |
| 1076 | 迷宫问题 | [Go](problems/1076) | 提高 |  |  |  |
| 1077 | 皇宫看守 | [Go](problems/1077) | 提高 |  |  |  |
| 1081 | 度的数量 | [Go](problems/1081) | 提高 |  |  |  |
| 1082 | 数字游戏 | [Go](problems/1082) | 提高 |  |  |  |
| 1083 | Windy数 | [Go](problems/1083) | 提高 |  |  |  |
| 1084 | 数字游戏II | [Go](problems/1084) | 提高 |  |  |  |
| 1085 | 不要62 | [Go](problems/1085) | 提高 |  |  |  |
| 1086 | 恨7不成妻 | [Go](problems/1086) | 提高 |  |  |  |
| 1087 | 修剪草坪 | [Go](problems/1087) | 提高 |  |  |  |
| 1088 | 旅行问题 | [Go](problems/1088) | 提高 |  |  |  |
| 1089 | 烽火传递 | [Go](problems/1089) | 提高 |  |  |  |
| 1090 | 绿色通道 | [Go](problems/1090) | 提高 |  |  |  |
| 1091 | 理想的正方形 | [Go](problems/1091) | 提高 |  |  |  |
| 1097 | 池塘计数 | [Go](problems/1097) | 提高 |  |  |  |
| 1098 | 城堡问题 | [Go](problems/1098) | 提高 |  |  |  |
| 1100 | 抓住那头牛 | [Go](problems/1100) | 提高 |  |  |  |
| 1106 | 山峰和山谷 | [Go](problems/1106) | 提高 |  |  |  |
| 1107 | 魔板 | [Go](problems/1107) | 提高 |  |  |  |
| 1112 | 迷宫 | [Go](problems/1112) | 提高 |  |  |  |
| 1113 | 红与黑 | [Go](problems/1113) | 提高 |  |  |  |
| 1116 | 马走日 | [Go](problems/1116) | 提高 |  |  |  |
| 1117 | 单词接龙 | [Go](problems/1117) | 提高 |  |  |  |
| 1118 | 分成互质组 | [Go](problems/1118) | 提高 |  |  |  |
| 1123 | 铲雪车 | [Go](problems/1123) | 提高 |  |  |  |
| 1124 | 骑马修栅栏 | [Go](problems/1124) | 提高 |  |  |  |
| 1125 | 牛的旅行 | [Go](problems/1125) | 提高 |  |  |  |
| 1126 | 最小花费 | [Go](problems/1126) | 提高 |  |  |  |
| 1127 | 香甜的黄油 | [Go](problems/1127) | 提高 |  |  |  |
| 1128 | 信使 | [Go](problems/1128) | 提高 |  |  |  |
| 1129 | 热浪 | [Go](problems/1129) | 提高 |  |  |  |
| 1131 | 拯救大兵瑞恩 | [Go](problems/1131) | 提高 |  |  |  |
| 1134 | 最短路计数 | [Go](problems/1134) | 提高 |  |  |  |
| 1135 | 新年好 | [Go](problems/1135) | 提高 |  |  |  |
| 1137 | 选择最佳线路 | [Go](problems/1137) | 提高 |  |  |  |
| 1140 | 最短网络 | [Go](problems/1140) | 提高 |  |  |  |
| 1141 | 局域网 | [Go](problems/1141) | 提高 |  |  |  |
| 1142 | 繁忙的都市 | [Go](problems/1142) | 提高 |  |  |  |
| 1143 | 联络员 | [Go](problems/1143) | 提高 |  |  |  |
| 1143 | 联络员 | [Go](problems/1143) | 提高 |  |  |  |
| 1144 | 连接格点 | [Go](problems/1144) | 提高 |  |  |  |
| 1145 | 北极通讯网络 | [Go](problems/1145) | 提高 |  |  |  |
| 1146 | 新的开始 | [Go](problems/1146) | 提高 |  |  |  |
| 1148 | 秘密的牛奶运输 | [Go](problems/1148) | 提高 |  |  |  |
| 1165 | 单词环 | [Go](problems/1165) | 提高 |  |  |  |
| 1169 | 糖果 | [Go](problems/1169) | 提高 |  |  |  |
| 1170 | 排队布局 | [Go](problems/1170) | 提高 |  |  |  |
| 1171 | 距离 | [Go](problems/1171) | 提高 |  |  |  |
| 1172 | 祖孙询问 | [Go](problems/1172) | 提高 |  |  |  |
| 1174 | 受欢迎的牛 | [Go](problems/1174) | 提高 |  |  |  |
| 1175 | 最大半连通子图 | [Go](problems/1175) | 提高 |  |  |  |
| 1183 | 电力 | [Go](problems/1183) | 提高 |  |  |  |
| 1184 | 欧拉回路 | [Go](problems/1184) | 提高 |  |  |  |
| 1185 | 单词游戏 | [Go](problems/1185) | 提高 |  |  |  |
| 1191 | 家谱树 | [Go](problems/1191) | 提高 |  |  |  |
| 1192 | 奖金 | [Go](problems/1192) | 提高 |  |  |  |
| 1250 | 格子游戏 | [Go](problems/1250) | 提高 |  |  |  |
| 1252 | 搭配购买 | [Go](problems/1252) | 提高 |  |  |  |
| 1273 | 天才的记忆 | [Go](problems/1273) | 提高 |  |  |  |
| 1275 | 最大数 | [Go](problems/1275) | 提高 |  |  |  |
| 1277 | 维护序列 | [Go](problems/1277) | 提高 |  |  |  |
| 1282 | 搜索关键词 | [Go](problems/1282) | 提高 |  |  |  |
| 1285 | 单词 | [Go](problems/1285) | 提高 |  |  |  |
| 1289 | 序列的第k个数 | [Go](problems/1289) | 提高 |  |  |  |
| 1290 | 越狱 | [Go](problems/1290) | 提高 |  |  |  |
| 1291 | 轻拍牛头 | [Go](problems/1291) | 提高 |  |  |  |
| 1292 | 哥德巴赫猜想 | [Go](problems/1292) | 提高 |  |  |  |
| 1293 | 夏洛克和他的女朋友 | [Go](problems/1293) | 提高 |  |  |  |
| 1294 | 樱花 | [Go](problems/1294) | 提高 |  |  |  |
| 1298 | 曹冲养猪 | [Go](problems/1298) | 提高 |  |  |  |
| 1303 | 斐波那契前n项和 | [Go](problems/1303) | 提高 |  |  |  |
| 1304 | 佳佳的斐波那契 | [Go](problems/1304) | 提高 |  |  |  |
| 1305 | GT考试 | [Go](problems/1305) | 提高 |  |  |  |
| 1307 | 牡牛和牝牛 | [Go](problems/1307) | 提高 |  |  |  |
| 1308 | 方程的解 | [Go](problems/1308) | 提高 |  |  |  |
| 1309 | 车的放置 | [Go](problems/1309) | 提高 |  |  |  |
| 1310 | 数三角形 | [Go](problems/1310) | 提高 |  |  |  |
| 1312 | 序列统计 | [Go](problems/1312) | 提高 |  |  |  |
| 1315 | 网格 | [Go](problems/1315) | 提高 |  |  |  |
| 1316 | 有趣的数列 | [Go](problems/1316) | 提高 |  |  |  |
| 1319 | 移棋子游戏 | [Go](problems/1319) | 提高 |  |  |  |
| 1321 | 取石子 | [Go](problems/1321) | 提高 |  |  |  |
| 1322 | 取石子游戏 | [Go](problems/1322) | 提高 |  |  |  |
| 0158 | 项链 | [Go](problems/0158) | 进阶 |  |  |  |
| 0169 | 数独2 | [Go](problems/0169) | 进阶 |  |  |  |
| 0182 | 破坏正方形 | [Go](problems/0182) | 进阶 |  |  |  |
| 0207 | 球形空间产生器 | [Go](problems/0207) | 进阶 |  |  |  |
| 0210 | 异或运算 | [Go](problems/0210) | 进阶 |  |  |  |
| 0221 | 龙哥的问题 | [Go](problems/0221) | 进阶 |  |  |  |
| 0243 | 一个简单的整数问题2 | [Go](problems/0243) | 进阶 |  |  |  |
| 0252 | 树 | [Go](problems/0252) | 进阶 |  |  |  |
| 0264 | 权值 | [Go](problems/0264) | 进阶 |  |  |  |
| 0304 | 诗人小G | [Go](problems/0304) | 进阶 |  |  |  |
| 0358 | 岛屿 | [Go](problems/0358) | 进阶 |  |  |  |
| 0359 | 创世纪 | [Go](problems/0359) | 进阶 |  |  |  |
| 0371 | 牧师约翰最忙碌的一天 | [Go](problems/0371) | 进阶 |  |  |  |
| 0381 | 有线电视网络 | [Go](problems/0381) | 进阶 |  |  |  |
| 0382 | K取方格数 | [Go](problems/0382) | 进阶 |  |  |  |
| 0516 | 神奇的幻方 | [Go](problems/0516) | 进阶 |  |  |  |
| 0918 | 软件包管理器 | [Go](problems/0918) | 进阶 |  |  |  |
| 0947 | 文本编辑器 | [Go](problems/0947) | 进阶 |  |  |  |
| 0950 | 郁闷的出纳员 | [Go](problems/0950) | 进阶 |  |  |  |
| 0955 | 维护数列 | [Go](problems/0955) | 进阶 |  |  |  |
| 0956 | 智慧珠游戏 | [Go](problems/0956) | 进阶 |  |  |  |
| 0961 | 最大获利 | [Go](problems/0961) | 进阶 |  |  |  |
| 0969 | 志愿者招募 | [Go](problems/0969) | 进阶 |  |  |  |
| 0999 | 魔法森林 | [Go](problems/0999) | 进阶 |  |  |  |
| 1004 | 品酒大会 | [Go](problems/1004) | 进阶 |  |  |  |
| 1032 | 游戏 | [Go](problems/1032) | 进阶 |  |  |  |
| 1063 | 永无乡 | [Go](problems/1063) | 进阶 |  |  |  |
| 1067 | 精确覆盖问题 | [Go](problems/1067) | 进阶 |  |  |  |
| 1080 | 骑士 | [Go](problems/1080) | 进阶 |  |  |  |
| 1283 | 玄武密码 | [Go](problems/1283) | 进阶 |  |  |  |
| 1358 | 约数个数和 | [Go](problems/1358) | 进阶 |  |  |  |
| 1401 | 围住奶牛 | [Go](problems/1401) | 进阶 |  |  |  |
| 1412 | 邮政货车 | [Go](problems/1412) | 进阶 |  |  |  |
| 2119 | 最佳包裹 | [Go](problems/2119) | 进阶 |  |  |  |
| 2142 | 最小矩形覆盖 | [Go](problems/2142) | 进阶 |  |  |  |
| 2144 | 神奇游乐园 | [Go](problems/2144) | 进阶 |  |  |  |
| 2154 | 梦幻布丁 | [Go](problems/2154) | 进阶 |  |  |  |
| 2171 | EK求最大流 | [Go](problems/2171) | 进阶 |  |  |  |
| 2172 | Dinic/ISAP求最大流 | [Go](problems/2172) | 进阶 |  |
| 2173 | Dinic/ISAP求最小割 | [Go](problems/2173) | 进阶 |  |
| 2174 | 费用流 | [Go](problems/2174) | 进阶 |  |  |  |
| 2175 | 飞行员配对方案问题 | [Go](problems/2175) | 进阶 |  |  |  |
| 2176 | 太空飞行计划问题 | [Go](problems/2176) | 进阶 |  |  |  |
| 2179 | 圆桌问题 | [Go](problems/2179) | 进阶 |  |  |  |
| 2180 | 最长递增子序列问题 | [Go](problems/2180) | 进阶 |  |  |  |
| 2184 | 餐巾计划问题 | [Go](problems/2184) | 进阶 |  |  |  |
| 2187 | 星际转移问题 | [Go](problems/2187) | 进阶 |  |  |  |
| 2188 | 无源汇上下界可行流 | [Go](problems/2188) | 进阶 |  |  |  |
| 2189 | 有源汇上下界最大流 | [Go](problems/2189) | 进阶 |  |  |  |
| 2190 | 有源汇上下界最小流 | [Go](problems/2190) | 进阶 |  |  |  |
| 2191 | 数字梯形问题 | [Go](problems/2191) | 进阶 |  |  |  |
| 2192 | 运输问题 | [Go](problems/2192) | 进阶 |  |  |  |
| 2193 | 分配问题 | [Go](problems/2193) | 进阶 |  |  |  |
| 2194 | 负载平衡问题 | [Go](problems/2194) | 进阶 |  |  |  |
| 2195 | 深海机器人问题 | [Go](problems/2195) | 进阶 |  |  |  |
| 2199 | 骑士共存问题 | [Go](problems/2199) | 进阶 |  |  |  |
| 2226 | 开店 | [Go](problems/2226) | 进阶 |  |  |  |
| 2234 | 多源汇最大流 | [Go](problems/2234) | 进阶 |  |  |  |
| 2236 | 伊基的故事I-道路重建 | [Go](problems/2236-伊基的故事I) | 进阶 |  |  |  |
| 2237 | 猪 | [Go](problems/2237) | 进阶 |  |  |  |
| 2240 | 餐饮 | [Go](problems/2240) | 进阶 |  |  |  |
| 2268 | 时态同步 | [Go](problems/2268) | 进阶 |  |  |  |
| 2277 | 秘密挤奶机 | [Go](problems/2277) | 进阶 |  |  |  |
| 2278 | 企鹅游行 | [Go](problems/2278) | 进阶 |  |  |  |
| 2279 | 网络战争 | [Go](problems/2279) | 进阶 |  |  |  |
| 2280 | 最优标号 | [Go](problems/2280) | 进阶 |  |  |  |
| 2306 | K大数查询 | [Go](problems/2306) | 进阶 |  |  |  |
| 2324 | 生活的艰辛 | [Go](problems/2324) | 进阶 |  |  |  |
| 2325 | 有向图破坏 | [Go](problems/2325) | 进阶 |  |  |  |
| 2326 | 王者之剑 | [Go](problems/2326) | 进阶 |  |  |  |
| 2402 | 2-SAT问题 | [Go](problems/2402-2) | 进阶 |  |  |  |
| 2417 | 指挥网络 | [Go](problems/2417) | 进阶 |  |  |  |
| 2418 | 光之大陆 | [Go](problems/2418) | 进阶 |  |  |  |
| 2419 | prufer序列 | [Go](problems/2419) | 进阶 |  |  |  |
| 2424 | 保龄球 | [Go](problems/2424) | 进阶 |  |  |  |
| 2437 | Splay | [Go](problems/2437) | 进阶 |  |  |  |
| 2476 | 树套树 | [Go](problems/2476) | 进阶 |  |  |  |
| 2488 | 树套树-简单版 | [Go](problems/2488-树套树) | 进阶 |  |  |  |
| 2492 | HH的项链 | [Go](problems/2492) | 进阶 |  |  |  |
| 2521 | 数颜色 | [Go](problems/2521) | 进阶 |  |  |  |
| 2523 | 历史研究 | [Go](problems/2523) | 进阶 |  |  |  |
| 2526 | 随机数生成器 | [Go](problems/2526) | 进阶 |  |  |  |
| 2534 | 树上计数2 | [Go](problems/2534) | 进阶 |  |  |  |
| 2535 | 二次离线莫队 | [Go](problems/2535) | 进阶 |  |  |  |
| 2539 | 动态树 | [Go](problems/2539) | 进阶 |  |  |  |
| 2568 | 树链剖分 | [Go](problems/2568) | 进阶 |  |  |  |
| 2572 | 生成魔咒 | [Go](problems/2572) | 进阶 |  |  |  |
| 2644 | 地板 | [Go](problems/2644) | 进阶 |  |  |  |
| 2680 | 均分数据 | [Go](problems/2680) | 进阶 |  |  |  |
| 2702 | problemb | [Go](problems/2702) | 进阶 |  |  |  |
| 2713 | 重复覆盖问题 | [Go](problems/2713) | 进阶 |  |  |  |
| 2714 | 左偏树 | [Go](problems/2714) | 进阶 |  |  |  |
| 2715 | 后缀数组 | [Go](problems/2715) | 进阶 |  |  |  |
| 2721 | K-单调 | [Go](problems/2721-K) | 进阶 |  |  |  |
| 2724 | 雷达 | [Go](problems/2724) | 进阶 |  |  |  |
| 2725 | 数字序列 | [Go](problems/2725) | 进阶 |  |  |  |
| 2752 | 仙人掌图 | [Go](problems/2752) | 进阶 |  |  |  |
| 2766 | 后缀自动机 | [Go](problems/2766) | 进阶 |  |  |  |
| 2785 | 信号增幅仪 | [Go](problems/2785) | 进阶 |  |  |  |
| 2801 | 三角形面积并 | [Go](problems/2801) | 进阶 |  |  |  |
| 2803 | 凸多边形 | [Go](problems/2803) | 进阶 |  |  |  |
| 2811 | 最长公共子串 | [Go](problems/2811) | 进阶 |  |  |  |
| 2815 | 三维偏序 | [Go](problems/2815) | 进阶 |  |  |  |
| 2819 | 动态逆序对 | [Go](problems/2819) | 进阶 |  |  |  |
| 2847 | 老C的任务 | [Go](problems/2847) | 进阶 |  |  |  |
| 2863 | 最短路 | [Go](problems/2863) | 进阶 |  |  |  |
| 2889 | 再探石子合并 | [Go](problems/2889) | 进阶 |  |  |  |
| 2934 | 插头DP | [Go](problems/2934) | 进阶 |  |  |  |
| 2935 | 信用卡凸包 | [Go](problems/2935) | 进阶 |  |  |  |
| 2938 | 周游世界 | [Go](problems/2938) | 进阶 |  |  |  |
| 2957 | 赛车 | [Go](problems/2957) | 进阶 |  |  |  |
| 2983 | 玩具 | [Go](problems/2983) | 进阶 |  |  |  |
| 2984 | 线段 | [Go](problems/2984) | 进阶 |  |  |  |
| 3020 | 建筑师 | [Go](problems/3020) | 进阶 |  |  |  |
| 3028 | 最小圆覆盖 | [Go](problems/3028) | 进阶 |  |  |  |
| 3034 | 望远镜 | [Go](problems/3034) | 进阶 |  |  |  |
| 3068 | 扫描线 | [Go](problems/3068) | 进阶 |  |  |  |
| 3069 | 圆的面积并 | [Go](problems/3069) | 进阶 |  |  |  |
| 3074 | 自适应辛普森积分 | [Go](problems/3074) | 进阶 |  |  |  |
| 3122 | 多项式乘法 | [Go](problems/3122) | 进阶 |  |  |  |
| 3123 | 高精度乘法II | [Go](problems/3123) | 进阶 |  |  |  |
| 3124 | BSGS | [Go](problems/3124) | 进阶 |  |  |  |
| 3125 | 扩展BSGS | [Go](problems/3125) | 进阶 |  |  |  |
| 3132 | 食物 | [Go](problems/3132) | 进阶 |  |  |  |
| 3133 | 串珠子 | [Go](problems/3133) | 进阶 |  |  |  |
| 3134 | 魔法手链 | [Go](problems/3134) | 进阶 |  |  |  |
| 3164 | 线性基 | [Go](problems/3164) | 进阶 |  |  |  |
| 3165 | 第一类斯特林数 | [Go](problems/3165) | 进阶 |  |  |  |
| 3166 | 第二类斯特林数 | [Go](problems/3166) | 进阶 |  |  |  |
| 3167 | 星星还是树 | [Go](problems/3167) | 进阶 |  |  |  |
| 3188 | manacher算法 | [Go](problems/3188) | 进阶 |  |  |  |
| 3189 | Lomsatgelral | [Go](problems/3189) | 进阶 |  |  |  |
| 3246 | 引水入城 | [Go](problems/3246) | 进阶 |  |  |  |

