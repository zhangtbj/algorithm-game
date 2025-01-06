package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	l := len(nums)

	if l == 1 {
		return nums[0]
	}

	rob1 := robWithoutCircle(nums[:l-1])
	rob2 := robWithoutCircle(nums[1:])

	return max(rob1, rob2)
}

func robWithoutCircle(nums []int) int {
	l := len(nums)

	var dp = make([]int, l+1) //step1 dp[i] 盗窃到第i个房间的最高金额
	//step3 init dp
	dp[0] = 0
	dp[1] = nums[0]

	//step4 loop
	for i := 2; i <= l; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1]) //step2 formula
	}

	//step5 debug
	fmt.Println(dp)

	return dp[l]
}
