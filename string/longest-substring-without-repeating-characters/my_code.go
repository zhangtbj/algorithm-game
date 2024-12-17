package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"

	fmt.Println(lengthOfLongestSubstringTest(s))
}

func lengthOfLongestSubstringTest(s string) int {
	var maxSize int
	for i := range s {
		var sMap = make(map[byte]bool)
		var subSize int
		for j := i; j < len(s); j++ {
			if !sMap[s[j]] {
				sMap[s[j]] = true
				subSize++
			} else {
				break
			}
		}
		maxSize = max(maxSize, subSize)
	}

	return maxSize
}