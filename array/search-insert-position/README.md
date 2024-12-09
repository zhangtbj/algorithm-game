# 搜索插入位置

本文档总结了代码实现的查找目标值插入位置的核心逻辑。

---

## 问题描述

给定一个有序数组 `nums` 和一个目标值 `target`，返回目标值将要插入的位置索引。如果目标值已经存在于数组中，返回其索引。

---

## 算法核心思路

### 1. 初始化指针
- 定义两个指针 `start` 和 `end`，分别指向数组的起始和末尾。

### 2. 二分查找
- 在 `start <= end` 的循环中执行以下操作：
    1. 计算中间索引 `mid = start + (end - start) >> 1`。
    2. **比较目标值和中间值 `nums[mid]`：**
        - **`target > nums[mid]`**：
            - 如果 `mid` 是数组最后一个索引，或者目标值小于等于 `nums[mid+1]`，返回 `mid + 1`，表示目标值应插入到当前值后面。
            - 否则，将搜索范围调整为右半部分：`start = mid + 1`。
        - **`target < nums[mid]`**：
            - 将搜索范围调整为左半部分：`end = mid - 1`。
        - **`target == nums[mid]`**：
            - 返回当前索引 `mid`，表示目标值已经存在于数组中。

### 3. 默认返回
- 如果循环结束时未找到目标值，返回 `0`，表示目标值应插入到数组的最前面。

---

## 时间复杂度

- **时间复杂度**：二分查找的时间复杂度为 $O(\log n)$，其中 $n$ 是数组的长度。
- **空间复杂度**：该算法使用常量级额外空间，空间复杂度为 $O(1)$。