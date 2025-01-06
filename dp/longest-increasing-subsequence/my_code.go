package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}

	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	l := len(nums)
	var maxLength = 1

	var dp = make([]int, l) //step1 dp[i]是以nums[i]为结尾的最长子序列长度
	//step3 init dp as 1
	for i := range dp {
		dp[i] = 1
	}

	//step4 loop
	for i := 0; i < l; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i]) //step2 formula
				maxLength = max(maxLength, dp[i])
			}
		}
	}

	//step5 debug
	fmt.Println(dp)

	return maxLength
}
