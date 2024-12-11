package main

import (
	"fmt"
)

func main() {
	var n = 5

	fmt.Println(generateMatrixTest(n))
}

func generateMatrixTest(n int) [][]int {
	var result = make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	row, col := 0, 0
	count := 1

	step := 0
	for count <= n*n {
		result[row][col] = count
		if row == 0+step && col < n-1-step {
			count++
			col++
			continue
		}
		if col == n-1-step && row < n-1-step {
			count++
			row++
			continue
		}
		if row == n-1-step && col > 0+step {
			count++
			col--
			continue
		}
		if col == 0+step && row > 1+step {
			count++
			row--
		}
		if col == 0+step && row == 1+step {
			step++
			continue
		}
		if count == n*n {
			break
		}
	}

	return result
}
