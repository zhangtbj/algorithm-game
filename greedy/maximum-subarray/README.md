# 53. 最大子序和
## 题目大意
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

## 视频讲解
https://www.bilibili.com/video/BV1aY4y1Z7ya?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
这一题可以用 DP 求解也可以不用 DP。
题目要求输出数组中某个区间内数字之和最大的那个值。dp[i] 表示 [0,i] 区间内各个子区间和的最大值，状态转移方程是 dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，dp[i] = nums[i] (dp[i-1] ≤ 0)。
