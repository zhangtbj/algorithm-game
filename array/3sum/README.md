# Three Sum 问题

## 问题描述

给定一个整数数组 `nums`，找出所有不重复的三元组，使得它们的和为零。三元组中的元素可以重复，但同样的三元组（即元素组合相同）只需要出现一次。

## 解决思路

### 1. **统计元素出现频次**

首先，通过一个哈希表（`numsMap`）统计数组中每个数字出现的次数。这样可以确保处理过程中能够方便地检查某个数字的出现次数，以便在后续查找过程中避免重复。

### 2. **去重与排序**

从 `numsMap` 中获取所有唯一的元素，并将这些元素存储在 `newNums` 数组中。然后对 `newNums` 数组进行排序，这样可以确保我们在后续遍历时可以利用排序后的顺序来简化查找。

### 3. **双重循环查找三元组**

- 外层循环：遍历排序后的数组 `newNums`，固定当前元素 `newNums[i]` 作为三元组的第一个数。
- 内层循环：对于每个 `newNums[i]`，我们用第二个指针 `newNums[j]` 从其后的元素开始寻找可能的配对。计算出第三个数 `c = -(newNums[i] + newNums[j])`，然后检查是否存在该元素并且满足条件。

### 4. **三元组生成**

根据以下几种情况来生成三元组：
- 当遇到特殊情况（如 0 0 0 三元组），需要特殊处理并添加到结果中。
- 如果发现 `newNums[i]` 与 `newNums[j]` 配对能够生成合法的三元组，并且该三元组没有重复，则将其添加到结果中。
- 通过检查是否存在第三个数 `c`，来确保生成合法的三元组。

### 5. **去重**

为了避免重复的三元组，需要确保同样的三元组（即元素顺序相同）只出现一次。通过判断 `newNums[i]` 和 `newNums[j]` 是否满足一定的出现次数，来避免重复处理相同的数值组合。

## 时间复杂度

- **时间复杂度**：
    - 统计频率的时间复杂度是 `O(n)`，其中 `n` 是数组的长度。
    - 数组排序的时间复杂度是 `O(k log k)`，其中 `k` 是去重后的数组长度，最坏情况下 `k = n`。
    - 双重循环遍历的时间复杂度是 `O(k^2)`，最坏情况下是 `O(n^2)`。
      因此，整体时间复杂度为 `O(n^2)`。

- **空间复杂度**：
    - 使用了一个哈希表 `numsMap` 来统计频率，空间复杂度为 `O(n)`。
    - 使用了一个数组 `newNums` 来存储去重后的元素，空间复杂度也是 `O(n)`。

## 使用场景

- 适用于需要寻找数组中符合某种条件的所有三元组的场景，尤其是在数组中的元素可能存在重复的情况下。
- 可以用于解决如“三数之和”这类典型的组合问题，或者是其他需要多元素配对和去重的算法题。

## 结论

通过统计元素频率、排序以及双重循环查找配对的方式，该算法能够高效地解决三数之和问题，避免了重复三元组的产生，且整体时间复杂度为 `O(n^2)`，适合大多数实际应用场景。