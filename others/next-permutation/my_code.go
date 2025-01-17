package main

import (
	"sort"
)

func main() {
	nums := []int{5, 1, 1}

	nextPermutation(nums)
}

// 1. 在 尽可能靠右的低位 进行交换，需要 从后向前 查找
// 2. 将一个 尽可能小的「大数」 与前面的「小数」交换。比如 123465，下一个排列应该把 5 和 4 交换而不是把 6 和 4 交换
// 3. 将「大数」换到前面后，需要将「大数」后面的所有数 重置为升序，升序排列就是最小的排列。以 123465 为例：首先按照上一步，交换 5 和 4，得到 123564；然后需要将 5 之后的数重置为升序，得到 123546。显然 123546 比 123564 更小，123546 就是 123465 的下一个排列
func nextPermutation(nums []int) {
	l := len(nums)
	var start = l - 2
	var end = l - 1
	for i := l - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			start--
			end--
		} else {
			break
		}
	}

	if start >= 0 {
		for i := l - 1; i >= start+1; i-- {
			if nums[i] > nums[start] {
				nums[i], nums[start] = nums[start], nums[i]
				break
			}
		}
	}

	subNums := nums[end:]
	sort.Slice(subNums, func(i, j int) bool {
		return subNums[i] < subNums[j]
	})
}
