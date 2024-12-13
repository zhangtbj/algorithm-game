package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	k := 2

	fmt.Println(reverseStr(s, k))
}

func reverseStr(s string, k int) string {
	sArr := []byte(s)

	for i := 0; i < len(sArr)-1; i += 2 * k {
		if i+k-1 < len(sArr) {
			reverse(sArr, i, i+k-1)
		} else {
			reverse(sArr, i, len(sArr)-1)
		}
	}

	return string(sArr)
}

func reverse(s []byte, start int, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
