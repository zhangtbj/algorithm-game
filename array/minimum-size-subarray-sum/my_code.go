package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	var target = 213

	fmt.Println(minSubArrayLenTest(target, nums))
}

// 我写错了，是因为这个array不能sort而改变他的顺序的
func minSubArrayLenTest(target int, nums []int) int {
	sort.Ints(nums)

	var count int
	end := len(nums) - 1
	for end >= 0 {
		if target > 0 {
			target -= nums[end]
			end--
			count++
		} else {
			break
		}
	}

	if target > 0 {
		count = 0
	}
	return count
}
