package main

import "fmt"

func main() {
	prices := []int{1, 2, 3, 0, 2}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	l := len(prices)

	var dp = make([][]int, l) //step1 dp[i][0]是第i天持有，dp[i][1]是第i天不持有，类似打家劫舍
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	//step4 loop
	for i := 1; i < l; i++ {
		if i >= 2 {
			dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i]) //step2 formula
		} else {
			dp[i][0] = max(dp[i-1][0], -prices[i])
		}
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return dp[l-1][1]
}
