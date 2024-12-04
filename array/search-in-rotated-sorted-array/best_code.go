package main

import (
	"fmt"
)

func main() {
	//	var nums = []int{4, 5, 6, 7, 0, 1, 2}
	var nums = []int{3, 1}
	var target = 0

	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	if len(nums) == 0 {
		return -1
	}

	for start <= end {
		mid := start + (end-start)>>1

		if target == nums[mid] {
			return mid
		} else if target == nums[start] {
			return start
		} else if target == nums[end] {
			return end
		} else if nums[mid] > nums[start] {
			if target > nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[end] {
			if target < nums[end] && target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[start] == nums[mid] {
				start++
			}
			if nums[end] == nums[mid] {
				end--
			}
		}
	}

	return -1
}
