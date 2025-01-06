package main

import "fmt"

func main() {
	n := 12

	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	var dp = make([]int, n+1) //step1 dp[i]是和为i的完全平方数的最少数量
	dp[0] = 0                 //step3 init dp
	for i := 1; i <= n; i++ {
		dp[i] = 10005
	}

	//step4 loop
	for i := 1; i <= 100 && i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1) //step2 formula
		}
	}

	//step5 debug
	fmt.Println(dp)

	return dp[n]
}
