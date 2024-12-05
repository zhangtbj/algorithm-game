package main

import (
	"fmt"
)

func main() {
	var nums = []int{5, 7, 7, 8, 8, 10}
	var target = 8

	fmt.Println(searchRangeTest(nums, target))
}

func searchRangeTest(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	startPos, endPos := -1, -1

	for start <= end {
		if nums[start] < target {
			start++
		}
		if start < len(nums) && nums[start] == target {
			startPos = start
		}

		if nums[end] > target {
			end--
		}
		if end >= 0 && nums[end] == target {
			endPos = end
		}

		if end >= 0 && start < len(nums) && nums[start] == nums[end] {
			break
		}
	}

	return []int{startPos, endPos}
}
