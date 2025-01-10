package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"

	fmt.Println(isSubsequence2(s, t))
}

func isSubsequence2(s string, t string) bool {
	ls := len(s)
	lt := len(t)

	var cur int

	for i := 0; i < lt; i++ {
		if cur < ls && t[i] == s[cur] {
			cur++
		}
	}

	if cur == ls {
		return true
	}
	return false
}
