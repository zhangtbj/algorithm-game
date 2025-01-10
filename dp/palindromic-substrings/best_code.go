package main

import (
	"fmt"
)

func main() {
	s := "abc"

	fmt.Println(countSubstrings(s))
}

func countSubstrings(s string) int {
	ls := len(s)
	var result int

	var dp = make([][]bool, ls) //step1 dp[i][j]是以i开始j结束的字符串是否是回文子串
	//step3 init dp
	for i := range dp {
		dp[i] = make([]bool, ls)
	}

	//step4 loop
	for i := ls - 1; i >= 0; i-- {
		for j := i; j < ls; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					result++
					dp[i][j] = true
				} else {
					if dp[i+1][j-1] {
						dp[i][j] = true //step2 formula
						result++
					}
				}
			}
		}
	}

	//step5 debug
	// for i := range dp {
	// 	fmt.Println(dp[i])
	// }

	return result
}
