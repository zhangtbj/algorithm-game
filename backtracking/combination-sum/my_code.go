package main

import (
	"fmt"
	"sort"
)

var result [][]int
var path []int

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	combinationSum(candidates, target)
	fmt.Println(result)
}

func combinationSum(candidates []int, target int) [][]int {
	result = [][]int{}
	path = []int{}
	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0)
	return result
}

func backtracking(candidates []int, target int, sum int, startIndex int) {
	if sum > target {
		return
	}

	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)

		return
	}

	for i := startIndex; i < len(candidates) && sum < target; i++ {
		sum += candidates[i]
		path = append(path, candidates[i])
		backtracking(candidates, target, sum, i)
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
