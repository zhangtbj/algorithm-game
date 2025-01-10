package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}

	fmt.Println(findLength2(nums1, nums2))
}

func findLength2(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	var count int

	var dp = make([][]int, l1)
	for i := range dp {
		dp[i] = make([]int, l2)
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = 1
			}
		}
	}

	for i := 0; i < l1; i++ {
		var subCount int
		for j := 0; j < l2; j++ {
			if i+j < l1 && dp[i+j][j] == 1 {
				subCount++
			} else {
				count = max(count, subCount)
				subCount = 0
			}
		}
		count = max(count, subCount)
	}

	for i := 1; i < l2; i++ {
		var subCount int
		for j := 0; j < l1; j++ {
			if i+j < l2 && dp[j][i+j] == 1 {
				subCount++
			} else {
				count = max(count, subCount)
				subCount = 0
			}
		}
		count = max(count, subCount)
	}

	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return count
}
