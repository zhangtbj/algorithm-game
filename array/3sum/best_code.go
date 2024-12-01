package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	var numsMap = make(map[int]int)
	for i := range nums {
		numsMap[nums[i]]++
	}

	var newNums []int
	for v := range numsMap {
		newNums = append(newNums, v)
	}

	sort.Ints(newNums)
	//fmt.Println(newNums)

	for i := range newNums {
		if newNums[i] == 0 && numsMap[0] >= 3 {
			result = append(result, []int{0, 0, 0})
		}
		for j := i + 1; j < len(newNums); j++ {
			if newNums[i]*2+newNums[j] == 0 && numsMap[newNums[i]] > 1 {
				result = append(result, []int{newNums[i], newNums[i], newNums[j]})
			}
			if newNums[i]+newNums[j]*2 == 0 && numsMap[newNums[j]] > 1 {
				result = append(result, []int{newNums[i], newNums[j], newNums[j]})
			}
			c := 0 - newNums[i] - newNums[j]
			if c > newNums[j] && numsMap[c] > 0 {
				result = append(result, []int{newNums[i], newNums[j], c})
			}
		}
	}
	return result
}
