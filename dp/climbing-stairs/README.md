# 70. 爬楼梯

## 题目大意
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？注意：给定 n 是一个正整数

## 视频讲解
https://www.bilibili.com/video/BV17h411h7UH?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
简单的 DP，经典的爬楼梯问题。一个楼梯可以由 n-1 和 n-2 的楼梯爬上来。
这一题求解的值就是斐波那契数列。
