package main

import (
	"fmt"
	"sort"
)

var result [][]int
var path []int
var used []bool

func main() {
	nums := []int{1, 1, 2}
	permuteUnique(nums)
	fmt.Println(result)
}

func permuteUnique(nums []int) [][]int {
	result = [][]int{}
	path = []int{}
	used = make([]bool, len(nums))
	sort.Ints(nums)
	backtracking(nums, path, used)
	return result
}

func backtracking(nums []int, path []int, used []bool) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
			continue
		}
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			backtracking(nums, path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
}
