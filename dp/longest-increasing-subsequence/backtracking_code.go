package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLIS2(nums))
}

var path []int
var count int

func lengthOfLIS2(nums []int) int {
	path = []int{}
	count = 0

	if len(nums) < 2 {
		return len(nums)
	}

	backtracking(nums, path, 0)
	return count
}

func backtracking(nums []int, path []int, startIndex int) {
	lp := len(path)
	if lp >= 2 {
		if path[lp-1] <= path[lp-2] {
			count = max(count, lp-1)
			return
		}
	}
	if startIndex == len(nums) {
		count = max(count, lp)
		return
	}

	for i := startIndex; i < len(nums); i++ {
		path = append(path, nums[i])
		//fmt.Println(path)
		backtracking(nums, path, i+1)
		path = path[:len(path)-1]
	}
}
