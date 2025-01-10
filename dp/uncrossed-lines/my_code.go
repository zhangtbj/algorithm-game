package main

import "fmt"

func main() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}

	fmt.Println(maxUncrossedLines(nums1, nums2))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)

	var dp = make([][]int, l1+1) //step1 dp[i][j]是nums1为[0,i-1]，nums2为[0,j-1]的最大连接数
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	//step4 loop
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 //step2 formula
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return dp[l1][l2]
}
