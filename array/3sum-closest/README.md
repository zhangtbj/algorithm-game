# 三数之和最接近 (`threeSumClosest`)

## **题目描述**

给定一个包含 `n` 个整数的数组 `nums` 和一个目标值 `target`，从数组中找到三个整数，使它们的和与 `target` 最接近，返回这三个数的和。假定每组输入只存在一个解。

---

## **解法思路**

### **1. 排序数组**
首先对数组进行排序，方便后续使用双指针法缩小范围。

### **2. 遍历数组**
- 使用一个外层循环固定第一个数 `nums[i]`。
- 内部使用双指针（`j` 和 `k`）寻找其他两个数的组合。

### **3. 双指针法**
- 初始化两个指针：
    - `j` 指向当前固定数之后的起始位置。
    - `k` 指向数组末尾。
- 计算三数之和 `sum = nums[i] + nums[j] + nums[k]`。
- 比较 `sum` 和 `target` 的差值，更新最接近的结果 `res`。
    - 如果 `sum == target`，直接返回结果。
    - 如果 `sum > target`，说明和偏大，移动右指针 `k--`。
    - 如果 `sum < target`，说明和偏小，移动左指针 `j++`。

### **4. 返回结果**
遍历完成后返回最接近 `target` 的三数之和。