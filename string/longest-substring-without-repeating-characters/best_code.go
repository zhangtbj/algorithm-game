package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"

	fmt.Println(lengthOfLongestSubstring(s))
}

// pwwkew
func lengthOfLongestSubstring(s string) int {
	var sMap = make(map[byte]bool)
	var maxSize int
	var end int

	for start := 0; start < len(s); start++ {
		if start > 0 {
			delete(sMap, s[start-1])
		}

		for end < len(s) && !sMap[s[end]] {
			sMap[s[end]] = true
			end++
		}
		maxSize = max(maxSize, end-start)
		//fmt.Printf("str is %s\n\n", s[start:end])
	}

	return maxSize
}
