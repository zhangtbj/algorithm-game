package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}

	fmt.Println(findLengthOfLCIS(nums))
}

func findLengthOfLCIS(nums []int) int {
	l := len(nums)
	var maxLength = 1

	var dp = make([]int, l) //step1 dp[i]是以nums[i]为尾的最长子序列
	//step3 init dp
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1 //step2 formula
			maxLength = max(maxLength, dp[i])
		}
	}

	//step5 debug
	//fmt.Println(dp)

	return maxLength
}
