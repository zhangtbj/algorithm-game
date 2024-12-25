package main

import "fmt"

var result [][]int
var path []int

func main() {
	n, k := 4, 2
	combinationSum3(n, k)
	fmt.Println(result)
}

func combinationSum3(k int, n int) [][]int {
	result = [][]int{}
	path = []int{}
	backtracking(k, n, 1, 0)
	return result
}

func backtracking(k, n, startIndex, sum int) {
	//fmt.Printf("found len(path): %d, k: %d, sum %d\n", len(path), k, sum)
	if sum > n {
		return
	}

	if len(path) == k {
		// sum := 0
		// for i := range path {
		//     sum += path[i]
		// }
		fmt.Printf("found sum: %d\n")
		if sum == n {
			fmt.Println("found one")
			tmp := make([]int, k)
			copy(tmp, path)
			result = append(result, tmp)
		}
		return
	}

	for i := startIndex; i <= 9; i++ {
		if 9-startIndex+1 < k-len(path) {
			break
		}

		// sum := 0
		// for i := range path {
		//     sum += path[i]
		// }
		// if sum > n {
		//     break
		// }
		sum += i
		path = append(path, i)
		//fmt.Printf("sum: %d, path: %i\n", sum, path)
		backtracking(k, n, i+1, sum)
		sum -= i
		path = path[:len(path)-1]
	}
}
