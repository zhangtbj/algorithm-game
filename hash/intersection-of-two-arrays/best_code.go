package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(nums1, nums2))
}

// The solution of array
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	var nums = make([]int, 1005)

	for i := range nums1 {
		nums[nums1[i]] = 1
	}

	for i := range nums2 {
		if nums[nums2[i]] == 1 {
			nums[nums2[i]] = 2
		}
	}

	for i := range nums {
		if nums[i] > 1 {
			result = append(result, i)
		}
	}

	return result
}
