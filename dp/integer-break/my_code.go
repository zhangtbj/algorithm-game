package main

import "fmt"

func main() {
	n := 10

	fmt.Println(integerBreak(n))
}

func integerBreak(n int) int {
	dp := make([]int, n+2) //step1 dp for max 乘积 of each num

	//step3 init dp
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp[2], dp[3] = 1, 2

	//step4 loop
	for i := 4; i <= n; i++ {
		for j := 2; j <= i/2; j++ {
			dp[i] = max(dp[i], max(dp[j], j)*max(dp[i-j], i-j)) // step2 formula
		}
	}

	//step5 debug
	fmt.Println(dp)

	return dp[n]
}
