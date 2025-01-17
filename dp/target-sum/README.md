# 494. 目标和

## 题目大意
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在有两个符号 + 和 -。对于数组中的任意一个整数，可以从 + 或 - 中选择一个符号添加在前面。返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

## 视频讲解
https://www.bilibili.com/video/BV1o8411j73x?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
给出一个数组，要求在这个数组里面的每个元素前面加上 + 或者 - 号，最终总和等于 S。问有多少种不同的方法。

这一题可以用 DP 和 DFS 解答。DFS 方法就不比较暴力简单了。见代码。这里分析一下 DP 的做法。题目要求在数组元素前加上 + 或者 - 号，其实相当于把数组分成了 2 组，一组全部都加 + 号，一组都加 - 号。记 + 号的一组 P ，记 - 号的一组 N，那么可以推出以下的关系。
sum(P) - sum(N) = target
sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
2 * sum(P) = target + sum(nums)
等号两边都加上 sum(N) + sum(P)，于是可以得到结果 2 * sum(P) = target + sum(nums)，那么这道题就转换成了，能否在数组中找到这样一个集合，和等于 (target + sum(nums)) / 2。那么这题就转化为了第 416 题了。dp[i] 中存储的是能使和为 i 的方法个数。

如果和不是偶数，即不能被 2 整除，那说明找不到满足题目要求的解了，直接输出 0 。

