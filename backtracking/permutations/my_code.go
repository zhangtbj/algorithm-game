package main

import (
	"fmt"
)

var result [][]int
var path []int
var used []bool

func main() {
	nums := []int{1, 2, 3}
	permute(nums)
	fmt.Println(result)
}

func permute(nums []int) [][]int {
	result = [][]int{}
	path = []int{}
	used = make([]bool, len(nums))
	backtracking(nums, path, used)
	return result
}

func backtracking(nums []int, path []int, used []bool) {
	//fmt.Println(path)
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
		return
	}

	// if len(path) >= len(nums) {
	//     return
	// }

	for i := 0; i < len(nums); i++ {
		if !used[nums[i]] {
			path = append(path, nums[i])
			used[nums[i]] = true
			backtracking(nums, path, used)
			path = path[:len(path)-1]
			used[nums[i]] = false
		}
	}
}
