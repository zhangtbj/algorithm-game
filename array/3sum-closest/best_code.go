package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var nums = []int{-1, 2, 1, -4}
	var target = 1

	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if int(math.Abs(float64(sum-target))) < diff {
					res, diff = sum, int(math.Abs(float64(sum-target)))
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}
