# 40. 组合总和 II

## 题目大意
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用一次。

## 视频讲解
https://www.bilibili.com/video/BV12V4y1V73A?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
题目要求出总和为 sum 的所有组合，组合需要去重。这一题是第 39 题的加强版，第 39 题中元素可以重复利用(重复元素可无限次使用)，这一题中元素只能有限次数的利用，因为存在重复元素，并且每个元素只能用一次(重复元素只能使用有限次)
这一题和第 47 题类似，只不过元素可以反复使用。
