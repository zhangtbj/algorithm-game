# 20. 有效的括号

## 题目大意
括号匹配问题。

## 视频讲解
https://www.bilibili.com/video/BV1AF411w78g?spm_id_from=333.788.player.switch&vd_source=f881def7ea7cf10e6fa73627efe940dd

## 解题思路
遇到左括号就进栈push，遇到右括号并且栈顶为与之对应的左括号，就把栈顶元素出栈。最后看栈里面还有没有其他元素，如果为空，即匹配。

需要注意，空字符串是满足括号匹配的，即输出 true。
