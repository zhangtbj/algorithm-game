package main

import "fmt"

func main() {
	test1 := "abcde"
	test2 := "ace"

	fmt.Println(longestCommonSubsequence(test1, test2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	var count int

	var dp = make([][]int, l1+1) //step1 dp[i][j]是以i为text1，j为text2的最长公共子序列
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
			count = max(count, dp[i][j])
		}
	}

	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return count
}
