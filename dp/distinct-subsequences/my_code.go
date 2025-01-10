package main

import "fmt"

func main() {
	s := "babgbag"
	t := "bag"

	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	ls := len(s)
	lt := len(t)

	var dp = make([][]int, ls+1) //step1 dp[i][j]：以i-1为结尾的s子序列中出现以j-1为结尾的t的个数为dp[i][j]
	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, lt+1)
		dp[i][0] = 1
	}

	//step4 loop
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lt; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] //step2 formula
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[ls][lt]
}
