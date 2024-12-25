package main

import (
	"fmt"
	"sort"
)

var result [][]int
var path []int
var used []bool

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	combinationSum2(candidates, target)
	fmt.Println(result)
}

func combinationSum2(candidates []int, target int) [][]int {
	result = [][]int{}
	path = []int{}
	used = make([]bool, len(candidates))
	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0, used)
	return result
}

func backtracking(candidates []int, target int, sum int, startIndex int, used []bool) {
	if sum > target {
		return
	}

	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		// Report timeout error
		// var pathStr string
		// for i := range tmp {
		//     pathStr += strconv.Itoa(tmp[i])
		// }
		// if !resultMap[pathStr] {
		//     resultMap[pathStr] = true
		//     result = append(result, tmp)
		// }
		result = append(result, tmp)

		return
	}

	for i := startIndex; i < len(candidates) && sum < target; i++ {
		// used is for same node check
		if i > 0 && candidates[i-1] == candidates[i] && !used[i-1] {
			continue
		}
		path = append(path, candidates[i])
		sum += candidates[i]
		used[i] = true
		backtracking(candidates, target, sum, i+1, used)
		sum -= candidates[i]
		path = path[:len(path)-1]
		used[i] = false
	}
}
