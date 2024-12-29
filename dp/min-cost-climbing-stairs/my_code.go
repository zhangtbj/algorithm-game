package main

import "fmt"

func main() {
	cost := []int{10, 15, 20}

	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	cost = append(cost, 0)
	dp := make([]int, l+1) //step1 dp为每一台阶所需cost

	dp[0] = cost[0] //step3, dp初始话
	dp[1] = cost[1]
	dp[l] = 0

	for i := 2; i <= l; i++ { //step4 遍历顺序
		dp[i] = cost[i] + min(dp[i-1], dp[i-2]) //step2 递推公式
	}

	fmt.Println(dp) // step5 print

	return dp[l]
}
