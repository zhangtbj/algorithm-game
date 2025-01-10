package main

import (
	"fmt"
)

func main() {
	s := "abc"

	fmt.Println(countSubstrings2(s))
}

func countSubstrings2(s string) int {
	ls := len(s)

	var dp = make([]int, ls) //step1 dp[i]是[0,i-i]的字符串的回文数目
	//step3 default is 1
	dp[0] = 1

	//step4 loop
	for i := 1; i < ls; i++ {
		var subCount int
		for j := 0; j < i; j++ {
			if isValid(s[j : i+1]) {
				subCount++
			}
		}
		dp[i] = dp[i-1] + subCount + 1 //step2 formula
	}

	//step5 debug
	fmt.Println(dp)

	return dp[ls-1]
}

func isValid(s string) bool {
	start, end := 0, len(s)-1
	for len(s) > 0 && start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}
	return true
}
