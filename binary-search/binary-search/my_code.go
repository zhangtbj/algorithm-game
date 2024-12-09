package main

import "fmt"

func main() {
	var nums = []int{-1, 0, 3, 5, 9, 12}
	var target = 9

	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)>>1

		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
