package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	l := len(prices)

	//step1 dp[i][0]第i天第一次持有股票的最大收益，dp[i][1]第i天第一次不持有股票的最大收益
	//      dp[i][2]第i天第二次持有股票的最大收益，dp[i][3]第i天第二次不持有股票的最大收益
	var dp = make([][]int, l)
	//step3 default for the first stock
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = -prices[0]
	dp[0][3] = 0

	//step4 loop
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i]) //step2 formula
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[l-1][3]
}
