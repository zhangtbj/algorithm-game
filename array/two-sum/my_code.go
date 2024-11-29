package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	fmt.Println(twoSumTest(nums, target))
}

func twoSumTest(nums []int, target int) []int {
	var nums2 []int
	var numsMap = make(map[int][]int)
	var nums3 []int

	for i := range nums {
		nums2 = append(nums2, target-nums[i])
		numsMap[nums[i]] = append(numsMap[nums[i]], i)
	}

	for i := range nums2 {
		if len(numsMap[nums2[i]]) > 0 {
			a := numsMap[nums2[i]][0]
			b := numsMap[target-nums2[i]][0]
			if nums2[i] == target/2 {
				if len(numsMap[nums2[i]]) > 1 {
					b = numsMap[nums2[i]][1]
				} else {
					continue
				}
			}

			if a < b {
				nums3 = []int{a, b}
			} else {
				nums3 = []int{b, a}
			}
			break
		}
	}

	return nums3
}
