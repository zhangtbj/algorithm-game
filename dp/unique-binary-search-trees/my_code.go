package main

import "fmt"

func main() {
	n := 5

	fmt.Println(numTrees(n))
}

func numTrees(n int) int {
	dp := make([]int, n+1) //step1 dp for the sum of tree types

	// step3 init dp
	dp[0] = 1

	// step4 loop
	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-j-1] //step2 formula
		}

	}

	fmt.Println(dp) //step5 debug

	return dp[n]
}
