package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit111(prices))
}

// Report timeout error
func maxProfit111(prices []int) int {
	var maxP, max1, max2 int
	l := len(prices)

	for i := 0; i < l-1; i++ {
		prices1 := prices[0 : i+1]
		prices2 := prices[i:]
		max1 = max(maxProfit1(prices1))
		max2 = max(maxProfit1(prices2))
		maxP = max(maxP, max1+max2)
	}

	return maxP
}

// 只买卖一次
func maxProfit1(prices []int) int {
	l := len(prices)

	var dp = make([][]int, l) //step1 dp[i][0]持有，dp[i][0]不持有
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]

	//step4 loop
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) //step2 formula
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return dp[l-1][1]
}
