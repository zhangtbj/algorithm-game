package main

import "fmt"

func main() {
	prices := []int{2, 4, 1}
	k := 2

	fmt.Println(maxProfit(k, prices))
}

func maxProfit(k int, prices []int) int {
	l := len(prices)

	//step1 dp[i][0]第i天第一次持有股票的最大收益，dp[i][1]第i天第一次不持有股票的最大收益
	//      dp[i][2]第i天第二次持有股票的最大收益，dp[i][3]第i天第二次不持有股票的最大收益
	var dp = make([][]int, l)
	//step3 default for the first stock
	for i := range dp {
		dp[i] = make([]int, k*2)
	}
	for i := 0; i < k*2; i++ {
		if i%2 == 0 {
			dp[0][i] = -prices[0]
		} else {
			dp[0][i] = 0
		}
	}

	//step4 loop
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i]) //step2 formula
		for j := 1; j < k*2; j++ {
			if j%2 == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			}
		}
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[l-1][k*2-1]
}
