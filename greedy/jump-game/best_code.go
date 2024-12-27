package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 3, 1, 1, 4}

	fmt.Println(canJump(nums))
}

func canJump(nums []int) bool {
	var path = 0

	if len(nums) == 1 {
		return true
	}

	for i := 0; i <= path; i++ {
		path = max(i+nums[i], path)
		if path >= len(nums)-1 {
			return true
		}
	}
	return false
}
