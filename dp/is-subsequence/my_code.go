package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"

	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)

	var dp = make([][]int, ls+1) //step1 dp[i][j]是[0,i-1]的s和[0,j-1]的t的最大公共子序列
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, lt+1)
	}

	//step4 loop
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lt; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1 //step2 formula
			} else {
				//dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				//s里没有多余的字符
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	//step5 debug
	// for i := range dp {
	//     fmt.Println(dp[i])
	// }

	if dp[ls][lt] == ls {
		return true
	}
	return false
}
