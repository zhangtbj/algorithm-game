package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{1, 0, -1, 0, -2, 2}
	var target = 0

	fmt.Println(fourSumTest(nums, target))
}

func fourSumTest(nums []int, target int) [][]int {
	var numsMap = make(map[int]int, len(nums))
	var newNums []int
	var result [][]int

	for i := range nums {
		numsMap[nums[i]]++
	}

	for v := range numsMap {
		newNums = append(newNums, v)
	}
	sort.Ints(newNums)
	//fmt.Println(numsMap)

	for i := 0; i < len(newNums); i++ {
		if newNums[i]*4 == target && numsMap[newNums[i]] >= 4 {
			result = append(result, []int{newNums[i], newNums[i], newNums[i], newNums[i]})
		}
		for j := i + 1; j < len(newNums); j++ {
			if newNums[i]*3+newNums[j] == target && numsMap[newNums[i]] >= 3 {
				result = append(result, []int{newNums[i], newNums[i], newNums[i], newNums[j]})
			}
			if newNums[i]*2+newNums[j]*2 == target && numsMap[newNums[i]] >= 2 && numsMap[newNums[j]] >= 2 {
				result = append(result, []int{newNums[i], newNums[i], newNums[j], newNums[j]})
			}
			if newNums[i]+newNums[j]*3 == target && numsMap[newNums[j]] >= 3 {
				result = append(result, []int{newNums[i], newNums[j], newNums[j], newNums[j]})
			}

			c := target - newNums[i]*2 - newNums[j]
			if c > newNums[j] && numsMap[newNums[i]] > 1 && numsMap[c] > 0 {
				result = append(result, []int{newNums[i], newNums[i], newNums[j], c})
			}
			d := target - newNums[i] - newNums[j]*2
			if d > newNums[j] && numsMap[newNums[j]] > 1 && numsMap[d] > 0 {
				result = append(result, []int{newNums[i], newNums[j], newNums[j], d})
			}

			e := target - newNums[i] - newNums[j]
			for k := j + 1; k < len(newNums); k++ {
				if e-newNums[k] > newNums[k] && numsMap[e-newNums[k]] > 0 {
					result = append(result, []int{newNums[i], newNums[j], newNums[k], e - newNums[k]})
				}
				if e-newNums[k] == newNums[k] && numsMap[e-newNums[k]] > 1 {
					result = append(result, []int{newNums[i], newNums[j], newNums[k], e - newNums[k]})
				}
			}
		}
	}

	return result
}
