package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 4, 7}

	fmt.Println(findLengthOfLCIS2(nums))
}

func findLengthOfLCIS2(nums []int) int {
	l := len(nums)
	var maxLength = 1

	var dp = make([]int, l) //step1 dp[i]是以nums[i]为尾的最长子序列
	//step3 init dp
	for i := range dp {
		dp[i] = 1
	}

	start := 0
	for i := 0; i < l; i++ {
		for j := start; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i]) //step2 formula
				maxLength = max(maxLength, dp[i])
			} else {
				dp[i] = 1
				start = j + 1
				continue
			}
		}
	}

	//step5 debug
	//fmt.Println(dp)

	return maxLength
}
