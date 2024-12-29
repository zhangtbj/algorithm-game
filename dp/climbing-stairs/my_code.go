package main

import "fmt"

func main() {
	n := 4

	fmt.Println(climbStairs(n))
}

func climbStairs(n int) int {
	dp := make([]int, n+1) //step1

	dp[0] = 0
	dp[1] = 1
	if n == 1 {
		return 1
	}
	dp[2] = 2
	if n == 2 {
		return 2
	}
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] //step2
	}

	fmt.Println(n)

	return dp[n]
}
