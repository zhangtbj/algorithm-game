package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcabcabcabc"

	fmt.Println(repeatedSubstringPatternTest(s))
}

func repeatedSubstringPatternTest(s string) bool {
	ss := s + s
	ss = ss[1 : len(ss)-1]

	return strings.Contains(ss, s)
}
