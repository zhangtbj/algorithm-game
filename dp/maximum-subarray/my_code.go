package main

import "fmt"

func main() {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums1))
}

func maxSubArray(nums []int) int {
	l := len(nums)
	var maxSum = nums[0]

	var dp = make([]int, l) //step1 dp[i]是[0,i-1]个连续子数组的最大和
	dp[0] = nums[0]

	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		maxSum = max(maxSum, dp[i])
	}

	fmt.Println(dp)

	return maxSum
}
