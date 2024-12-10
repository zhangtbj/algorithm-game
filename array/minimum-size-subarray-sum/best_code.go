package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 3, 1, 2, 4, 3}
	var target = 7

	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	var sum, subL int
	count := len(nums) + 1
	i := 0
	for j := 0; j <= len(nums)-1; j++ {
		sum += nums[j]
		for sum >= target {
			subL = j - i + 1
			count = min(count, subL)
			sum -= nums[i]
			i++
		}
	}
	if count > len(nums) {
		count = 0
	}
	return count
}
