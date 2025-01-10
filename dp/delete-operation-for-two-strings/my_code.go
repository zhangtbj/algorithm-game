package main

import "fmt"

func main() {
	word1 := "sea"
	word2 := "eat"

	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)

	var dp = make([][]int, l1+1) //step1 dp[i][j]是使[0,i-1]的word1和[0,j-1]的word2相同所需要删除的最小步数
	//step3 init dp {
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}

	for i := 0; i <= l2; i++ {
		dp[0][i] = i
	}

	//step4 loop
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] //step2 formula
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[l1][l2]
}
