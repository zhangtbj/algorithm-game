package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 5

	fmt.Println(change(amount, coins))
}

func change(amount int, coins []int) int {
	l := len(coins)

	var dp = make([]int, amount+1) //step1 dp for complete pack

	//step3 default is 0
	dp[0] = 1

	//step4 loop
	for i := 0; i < l; i++ {
		for j := coins[i]; j <= amount; j++ {
			//fmt.Printf("dp[%d]: %d, dp[j-coins[i]: %d\n", j, dp[j], dp[j-coins[i]])
			dp[j] += dp[j-coins[i]] //step2 formula
		}
		//fmt.Println(dp)
	}

	//step5 debug
	//fmt.Println(dp)

	return dp[amount]
}
