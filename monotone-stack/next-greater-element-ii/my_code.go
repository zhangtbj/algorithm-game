package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1}

	fmt.Println(nextGreaterElements2(nums))
}

// 把两个数组拼起来的方式
func nextGreaterElements2(nums []int) []int {
	l := len(nums)
	var result = make([]int, l*2)
	for i := range result {
		result[i] = -1
	}
	newNums := nums
	newNums = append(newNums, nums...)

	var stack []int
	stack = append(stack, 0)

	for i := 1; i < l*2; i++ {
		if len(stack) > 0 && newNums[i] > newNums[stack[len(stack)-1]] {
			for len(stack) > 0 && newNums[i] > newNums[stack[len(stack)-1]] {
				result[stack[len(stack)-1]] = newNums[i]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
	}

	return result[:l]
}
