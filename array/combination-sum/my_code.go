package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{2, 3, 5}
	var target = 8

	fmt.Println(combinationSumTest(nums, target))
}

func combinationSumTest(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var test []int
	var result [][]int

	for i := 0; i < len(candidates); i++ {
		test = append(test, candidates[i])
		for j := 0; j < len(candidates); j++ {
			var sum int
			for sum < target {
				sum = 0
				for k := 0; k < len(test); k++ {
					sum += test[k]
				}
				if sum+candidates[j] < target {
					test = append(test, candidates[j])
				}
				if sum+candidates[j] == target {
					result = append(result, append(test, candidates[j]))
					break
				}
				if sum+candidates[j] > target {
					test = test[:len(test)-1]
					if test[0] != test[1] {
						break
					}
				}
			}

		}
	}

	return result
}
