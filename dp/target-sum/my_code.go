package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3

	fmt.Println(findTargetSumWays(nums, target))
}

func findTargetSumWays(nums []int, target int) int {
	var sum int
	var bigTarget int
	for i := range nums {
		sum += nums[i]
	}
	if (sum+target)%2 != 0 {
		return 0
	}
	if math.Abs(float64(target)) > float64(sum) {
		return 0
	}

	bigTarget = (sum + target) / 2

	l := len(nums)
	//fmt.Println(nums)

	//step1 dp as 10pack
	dp := make([]int, bigTarget+1)

	//step3 default 0 for dp init
	dp[0] = 1

	//step4 loop
	for i := 0; i < l; i++ {
		for j := bigTarget; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]] //step2 formula
			//fmt.Println(dp)
		}
	}

	//step5 debug
	fmt.Println()
	fmt.Println(dp)

	return dp[bigTarget]
}
