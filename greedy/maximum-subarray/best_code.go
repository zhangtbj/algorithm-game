package main

import (
	"fmt"
)

func main() {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var result = nums[0]
	var checkNum int

	if len(nums) == 1 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		checkNum += nums[i]
		if checkNum < 0 {
			checkNum = 0
			result = max(result, nums[i])
		} else {
			result = max(result, checkNum)
		}
	}

	return result
}
