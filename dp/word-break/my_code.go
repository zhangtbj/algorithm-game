package main

import "fmt"

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	l := len(s)
	var dp = make([]string, l+1) //step1 dp[i]是由wordDict组合的字符串
	//step3 default is empty
	dp[0] = ""

	//step4 loop
	for j := 1; j <= len(s); j++ {
		for i := 0; i < len(wordDict); i++ {
			if j-len(wordDict[i]) >= 0 && s[:j] == dp[j-len(wordDict[i])]+wordDict[i] {
				//fmt.Printf("j-len(wordDict[i]): %d, dp[j-len(wordDict[i])]: %s, s[:j-len(wordDict[i])]: %s, wordDict[i]: %s\n", j-len(wordDict[i]), dp[j-len(wordDict[i])], s[:j-len(wordDict[i])], wordDict[i])
				dp[j] = dp[j-len(wordDict[i])] + wordDict[i] //step2 dormula
				break
			}
		}
	}

	//step5 debug
	fmt.Println(dp)

	if dp[l] == s {
		return true
	}
	return false
}
