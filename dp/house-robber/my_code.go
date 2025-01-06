package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}

	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	l := len(nums)

	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}

	var dp = make([]int, l+1) //step1 dp[i]是偷到第i家时的最大值
	//step3 init dp
	dp[0] = 0
	dp[1] = nums[0]

	//step4 loop
	for i := 2; i <= l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1]) //step2 forumla
	}

	//step5 debug
	fmt.Println(dp)

	return dp[l]
}
