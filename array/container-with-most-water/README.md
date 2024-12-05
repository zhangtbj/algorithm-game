# 最大水容器问题 (Max Area)

该项目实现了经典的 **最大水容器问题**，给定一个整数数组 `height`，数组中的每个元素表示一个垂直的线段，线段的宽度为 1。目标是找到两个线段，能够形成的最大水容器的面积。

## 问题描述

给定一个整数数组 `height`，其中 `height[i]` 代表第 `i` 条线段的高度。找出两个线段，能够形成的最大面积，计算方式如下：

- 容器的宽度是两条线段之间的距离，记作 `high - low`。
- 容器的高度是两条线段的较短者，记作 `min(height[low], height[high])`。

所以，容器的面积为 `min(height[low], height[high]) * (high - low)`。

## 算法说明

该实现采用 **双指针** 方法，通过维护两个指针 `low` 和 `high`，分别指向数组的两端，逐步缩小范围并计算最大面积。算法的时间复杂度是 **O(n)**，其中 `n` 是数组 `height` 的长度。

- 如果 `height[low]` 小于 `height[high]`，则增加 `low` 指针，尝试寻找更高的线段。
- 否则，减少 `high` 指针，尝试寻找更高的线段。
- 在每次计算当前面积后，更新最大面积 `maxS`。