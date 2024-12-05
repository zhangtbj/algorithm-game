package main

import (
	"fmt"
)

func main() {
	var nums = []int{5, 7, 7, 8, 8, 10}
	var target = 8

	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

func searchFirstEqualElement(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			end = mid - 1
		}
	}
	return -1
}

func searchLastEqualElement(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			start = mid + 1
		}
	}
	return -1
}

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后一个小于等于 target 的元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
