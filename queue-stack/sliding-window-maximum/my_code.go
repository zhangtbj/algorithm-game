package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindowTest(nums, k))
}

func maxSlidingWindowTest(nums []int, k int) []int {
	size := len(nums) - k + 1
	var result = make([]int, size)

	for i := 0; i < size; i++ {
		maxV := -10005
		for j := i; j < i+k; j++ {
			if nums[j] >= maxV {
				maxV = nums[j]
			}
		}
		result[i] = maxV
	}
	return result
}
