package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2

	fmt.Println(maxProfit(prices, fee))
}

func maxProfit(prices []int, fee int) int {
	l := len(prices)

	var dp = make([][]int, l) //step1 dp[i][0]是第i天持有股票，dp[i][1]是第i天不持有股票抛掉手续费
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	//step4 loop
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]) //step2 formula
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return dp[l-1][1]
}
