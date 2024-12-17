# 226. 翻转二叉树

## 题目大意
"经典"的反转二叉树的问题。

## 视频讲解
https://www.bilibili.com/video/BV1sP4y1f7q7?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
还是用递归来解决，先递归调用反转根节点的左孩子，然后递归调用反转根节点的右孩子，然后左右交换根节点的左孩子和右孩子。
