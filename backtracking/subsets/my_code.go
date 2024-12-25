package main

import (
	"fmt"
	"sort"
)

var result [][]int
var path []int

func main() {
	nums := []int{1, 2, 3}
	subsets(nums)
	fmt.Println(result)
}

func subsets(nums []int) [][]int {
	result = [][]int{{}}
	path = []int{}
	sort.Ints(nums)
	backtracking(nums, path, 0)
	return result
}

func backtracking(nums []int, path []int, startIndex int) {
	if len(path) > 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
	}
	if startIndex == len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		backtracking(nums, path, i+1)
		path = path[:len(path)-1]
	}
}
