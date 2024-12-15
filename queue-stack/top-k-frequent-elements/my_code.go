package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) []int {
	numsMap := map[int]int{}
	var result []int

	for i := range nums {
		numsMap[nums[i]]++
	}

	for key, _ := range numsMap {
		result = append(result, key)
	}

	sort.Slice(result, func(i, j int) bool {
		return numsMap[result[i]] > numsMap[result[j]]
	})

	return result[:k]
}
