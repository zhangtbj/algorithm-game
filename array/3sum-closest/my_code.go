package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var nums = []int{-1, 2, 1, -4}
	var target = 1

	fmt.Println(threeSumClosestTest(nums, target))
}

func threeSumClosestTest(nums []int, target int) int {
	sort.Ints(nums)
	var minFV = 100000
	var minFR int

	for i := 0; i < len(nums)-2; i++ {
		start := i + 1
		end := len(nums) - 1
		minV := 100000
		var minR int
		for start < end {
			if minV > int(math.Abs(float64(nums[i]+nums[start]+nums[end]-target))) {
				minV = int(math.Abs(float64(nums[i] + nums[start] + nums[end] - target)))
				minR = nums[i] + nums[start] + nums[end]
			}
			if nums[i]+nums[start]+nums[end]-target > 0 {
				end--
			} else {
				start++
			}
		}
		if minFV > minV {
			minFV = minV
			minFR = minR
		}
	}
	return minFR
}
