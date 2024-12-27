package main

import (
	"fmt"
)

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	sum := 0

	if len(prices) == 1 {
		return sum
	}

	for i := 0; i < len(prices)-1; i++ {
		//if prices[i+1]-prices[i] > 0 {
		//	sum += prices[i+1] - prices[i]
		//}

		// 简单解法
		sum += max(prices[i+1]-prices[i], 0)
	}

	return sum
}
