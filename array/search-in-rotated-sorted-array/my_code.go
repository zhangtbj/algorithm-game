package main

import (
	"fmt"
)

func main() {
	//	var nums = []int{4, 5, 6, 7, 0, 1, 2}
	var nums = []int{8, 9, 2, 3, 4}
	var target = 9

	fmt.Println(searchTest(nums, target))
}

// O(n) is not good
//func searchTest(nums []int, target int) int {
//	var result = -1
//	start := 0
//	end := len(nums) - 1
//
//	for start <= end {
//		if target != nums[start] && target != nums[end] {
//			start++
//			end--
//		}
//
//		if target == nums[start] {
//			result = start
//			break
//		}
//		if target == nums[end] {
//			result = end
//			break
//		}
//
//	}
//
//	return result
//}

func searchTest(nums []int, target int) int {
	var result = -1
	start := 0
	end := len(nums) - 1
	if target == nums[end] {
		return end
	}

	for start < end {
		if target == nums[end] {
			result = end
			break
		}
		mid := (start + end + 1) / 2
		if target > nums[start] {
			if target < nums[mid] {
				end = mid - 1
			} else if target > nums[mid] {
				if nums[start] > nums[mid] {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else {
				result = mid
				break
			}
		} else if target < nums[start] {
			if target < nums[mid] {
				if nums[start] > nums[mid] {
					end = mid - 1
				} else {
					start = mid + 1
				}
			} else if target > nums[mid] {
				start = mid + 1
			} else {
				result = mid
				break
			}
		} else {
			result = start
			break
		}
	}

	return result
}
