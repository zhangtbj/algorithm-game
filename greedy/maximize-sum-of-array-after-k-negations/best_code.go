package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var nums = []int{3, -1, 0, 2}

	k := 3

	fmt.Println(largestSumAfterKNegations(nums, k))
}

func largestSumAfterKNegations(nums []int, k int) int {
	var sum int
	l := len(nums)
	//sort.Ints(nums)
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < l && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	//sort.Ints(nums)
	if k > 0 && k%2 != 0 {
		nums[l-1] = -nums[l-1]
	}

	for i := range nums {
		sum += nums[i]
	}

	return sum
}
