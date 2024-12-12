package main

import (
	"fmt"
)

func main() {
	var nums1 = []int{1, 2}
	var nums2 = []int{-2, -1}
	var nums3 = []int{-1, 2}
	var nums4 = []int{0, 2}

	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var sum1 = make(map[int]int)
	var count int

	for i := range nums1 {
		for j := range nums2 {
			sum1[nums1[i]+nums2[j]]++
		}
	}

	for i := range nums3 {
		for j := range nums4 {
			sum2 := nums3[i] + nums4[j]
			if _, ok := sum1[0-sum2]; ok {
				count += sum1[0-sum2]
			}
		}
	}

	return count
}
