package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 11

	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	l := len(coins)

	var dp = make([]int, amount+1) //step1 for complete pack

	//step3 init dp
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = 10005
	}

	//step4 loop
	for i := 0; i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1) //step2 formula
			//fmt.Println(dp)
		}
	}

	//step5 debug
	fmt.Println(dp)

	if dp[amount] == 10005 {
		return -1
	}
	return dp[amount]
}
