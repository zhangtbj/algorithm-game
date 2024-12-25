package main

import (
	"fmt"
	"sort"
)

var result [][]int
var path []int
var used []bool

func main() {
	nums := []int{1, 2, 3}
	subsetsWithDup(nums)
	fmt.Println(result)
}

func subsetsWithDup(nums []int) [][]int {
	result = [][]int{}
	path = []int{}
	used = make([]bool, len(nums))
	sort.Ints(nums)
	backtracking(nums, path, 0, used)
	return result
}

func backtracking(nums []int, path []int, startIndex int, used []bool) {
	if startIndex <= len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
	} else {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtracking(nums, path, i+1, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
