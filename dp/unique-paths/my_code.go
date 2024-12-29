package main

import "fmt"

func main() {
	m := 3
	n := 7

	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m) //step1 dp is count of each path

	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1

		if i == 0 {
			for j := 0; j < n; j++ {
				dp[0][j] = 1
			}
		}
	}

	//step4 loop
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] //step2 formula
		}
	}

	for i := range dp {
		fmt.Println(dp[i]) // step5 debug
	}

	return dp[m-1][n-1]
}
