package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 3, 5, 6}
	var target = 7

	fmt.Println(searchInsertTest(nums, target))
}

func searchInsertTest(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)>>1
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}

	return start
}
