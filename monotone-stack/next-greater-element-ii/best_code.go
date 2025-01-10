package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1}

	fmt.Println(nextGreaterElements(nums))
}

// 取模的方式
func nextGreaterElements(nums []int) []int {
	l := len(nums)
	var result = make([]int, l)
	for i := range result {
		result[i] = -1
	}

	var stack []int
	stack = append(stack, 0)

	for i := 1; i < l*2; i++ {
		if len(stack) > 0 && nums[i%l] > nums[stack[len(stack)-1]] {
			for len(stack) > 0 && nums[i%l] > nums[stack[len(stack)-1]] {
				result[stack[len(stack)-1]] = nums[i%l]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i%l)
	}

	return result[:l]
}
