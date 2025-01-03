package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	target := 4

	fmt.Println(combinationSum4(nums, target))
}

func combinationSum4(nums []int, target int) int {
	var dp = make([]int, target+1) //step1 dp for complete pack
	//step3 init dp
	dp[0] = 1

	//step4 loop
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]] //step2 formula
			}
		}
	}

	//step5 debug
	fmt.Println(dp)

	return dp[target]
}
