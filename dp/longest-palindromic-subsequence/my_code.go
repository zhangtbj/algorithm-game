package main

import "fmt"

func main() {
	s := "bbbab"

	fmt.Println(longestPalindromeSubseq(s))
}

func longestPalindromeSubseq(s string) int {
	ls := len(s)
	var count int

	var dp = make([][]int, ls) //step1 dp[i][j]是[i,j]的字符串是否改变成回文子序列的最大唱的
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, ls)
	}

	//step4 loop
	for i := ls - 1; i >= 0; i-- {
		for j := i; j < ls; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = j - i + 1
				} else {
					dp[i][j] = dp[i+1][j-1] + 2 //step2 formula
				}
				count = max(count, dp[i][j])
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	return count
}
