package main

import "fmt"

var result [][]int
var path []int

func main() {
	n, k := 4, 2
	combine(n, k)
	fmt.Println(result)
}

func combine(n int, k int) [][]int {
	result = [][]int{}
	path = []int{}
	backtracking(n, k, 1)
	return result
}

func backtracking(n int, k int, startIndex int) {
	if len(path) == k {
		// 切片是引用类型，直接修改值会导致result里所有引用都被修改
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, path)
		return
	}

	for i := startIndex; i <= n; i++ {
		// 减枝， n - i + 1 为 剩下的数字， k - len(path) 为 需要的数字
		if n-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		backtracking(n, k, i+1)
		path = path[:len(path)-1]
	}
}
