package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0}, {0, 1, 0}, {0, 0, 0},
	}

	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m) //step1 dp for count of path

	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 || (i > 0 && dp[i-1][0] == 0) {
			dp[i][0] = 0
		} else {
			dp[i][0] = 1
		}
		for j := 0; j < n; j++ {
			if obstacleGrid[0][j] == 1 || (j > 0 && dp[0][j-1] == 0) {
				dp[0][j] = 0
			} else {
				dp[0][j] = 1
			}
		}
	}

	//step4 loop
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] //step2 formula
			}
		}
	}

	for i := range dp {
		fmt.Println(dp[i]) //step5 debug
	}

	return dp[m-1][n-1]
}
