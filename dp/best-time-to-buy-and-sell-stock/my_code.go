package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	l := len(prices)
	var dp = make([][]int, l) //step1 dp[i][0]是第i天的持股最大收益，dp[i][1]是第i天的不持股最大收益
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	//step3 default is the first day deal
	dp[0][0] = -prices[0]

	//step4 loop
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[l-1][1]
}
