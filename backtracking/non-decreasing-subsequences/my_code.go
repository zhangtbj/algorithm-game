package main

import (
	"fmt"
)

var result [][]int
var path []int

//var used []bool

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}
	findSubsequences(nums)
	fmt.Println(result)
}

func findSubsequences(nums []int) [][]int {
	result = [][]int{}
	path = []int{}
	//used = make([]bool, len(nums))
	backtracking(nums, path, 0)
	return result
}

func backtracking(nums []int, path []int, startIndex int) {
	if startIndex <= len(nums) {
		if len(path) > 1 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}
	} else {
		return
	}

	// 初始化used字典，用以对同层元素去重
	used := make(map[int]bool)
	for i := startIndex; i < len(nums); i++ {
		// var skip bool
		// for j:=startIndex; j<i;j++ {
		//     if i > 0 && nums[j] == nums[i] && !used[j] {
		//         skip=true
		//         break
		//     }
		// }
		// if skip {
		//     continue
		// }

		if used[nums[i]] {
			continue
		}

		if len(path) > 0 && path[len(path)-1] > nums[i] {
			continue
		}

		path = append(path, nums[i])
		used[nums[i]] = true
		// fmt.Printf("path after append: ")
		// fmt.Println(path)
		backtracking(nums, path, i+1)
		//used[i] = false
		path = path[:len(path)-1]
		// fmt.Printf("path after clean: ")
		// fmt.Println(path)
	}

}
