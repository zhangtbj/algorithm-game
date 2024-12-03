package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{1, 1, 2}

	fmt.Println(removeDuplicatesTest(nums))
}

func removeDuplicatesTest(nums []int) int {
	numsMap := make(map[int]int, len(nums))
	var newNums []int

	for i := range nums {
		numsMap[nums[i]]++
	}

	for v := range numsMap {
		newNums = append(newNums, v)
	}

	sort.Ints(newNums)

	for i := range newNums {
		nums[i] = newNums[i]
	}

	return len(newNums)
}
