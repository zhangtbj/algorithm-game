package main

import (
	"fmt"
)

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var numsMap = make(map[int]int, len(nums2))
	for i := range nums2 {
		numsMap[nums2[i]] = -1
	}
	var stack []int
	var result = make([]int, len(nums1))
	stack = append(stack, nums2[0])

	for i := 1; i < len(nums2); i++ {
		if len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
				numsMap[stack[len(stack)-1]] = nums2[i]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, nums2[i])
	}

	for i := range nums1 {
		result[i] = numsMap[nums1[i]]
	}

	return result
}
