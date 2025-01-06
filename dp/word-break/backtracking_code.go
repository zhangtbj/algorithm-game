package main

import "fmt"

// Will report timeout, 只需要求是否存在，不用列出所有结果，所以backtracking不合适
func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	fmt.Println(wordBreak2(s, wordDict))
}

func wordBreak2(s string, wordDict []string) bool {
	cur := ""
	return backtracking(s, wordDict, cur)
}

func backtracking(s string, wordDict []string, cur string) bool {
	//fmt.Printf("s: %s, cur: %s\n", s, cur)
	lc := len(cur)
	ls := len(s)
	if lc == ls {
		if cur == s {
			return true
		} else {
			return false
		}
	} else if lc > ls {
		return false
	} else {
		if lc > 0 && cur != s[:lc] {
			return false
		}
	}

	for i := 0; i < len(wordDict); i++ {
		cur += wordDict[i]
		//fmt.Printf("cur add: %s\n", cur)
		if backtracking(s, wordDict, cur) {
			return true
		}
		cur = cur[:len(cur)-len(wordDict[i])]
	}
	return false
}
