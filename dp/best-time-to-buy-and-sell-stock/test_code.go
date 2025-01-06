package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit2(prices))
}

func maxProfit2(prices []int) int {
	l := len(prices)
	var dp = make([]int, l+1) //step1 dp[i]是第i天的最大收益
	//step3 default is 0
	minPrice := 10005

	//step4 loop
	for i := 1; i <= l; i++ {
		minPrice = min(prices[i-1], minPrice)
		dp[i] = max(dp[i-1], prices[i-1]-minPrice) //step2 formula
	}

	//step5 debug
	//fmt.Println(dp)

	return dp[l]
}
