# 78. 子集

## 题目大意
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。

## 视频讲解
https://www.bilibili.com/video/BV1U84y1q7Ci?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
找出一个集合中的所有子集，空集也算是子集。且数组中的数字不会出现重复。用 DFS 暴力枚举即可。
这一题和第 90 题，第 491 题类似，可以一起解答和复习。
