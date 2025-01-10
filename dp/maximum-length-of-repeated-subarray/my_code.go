package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}

	fmt.Println(findLength(nums1, nums2))
}

func findLength(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	var count int

	var dp = make([][]int, l1+1) //step1 dp[i][j]是以i为结尾的nums1和j为结尾的nums2的最长子数组
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	//step4 loop
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 //step2 formula
				count = max(count, dp[i][j])
			}
		}
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return count
}
