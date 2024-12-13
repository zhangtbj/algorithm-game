package main

import (
	"fmt"
)

func main() {
	s := "  hello world  "

	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	sArr := []byte(s)

	slow := 0
	for fast := 0; fast < len(sArr); fast++ {
		if sArr[fast] != ' ' {
			sArr[slow] = sArr[fast]
			slow++
		}

		if slow != 0 && sArr[fast] == ' ' && fast+1 < len(sArr) && sArr[fast+1] != ' ' {
			sArr[slow] = ' '
			slow++
		}
	}
	sArr = sArr[:slow]

	reverse(sArr)

	var pos int
	for i := range sArr {
		if sArr[i] == ' ' {
			reverse(sArr[pos:i])
			pos = i + 1
		}
		if i == len(sArr)-1 {
			reverse(sArr[pos : i+1])
		}
	}

	return string(sArr)
}

func reverse(s []byte) {
	start, end := 0, len(s)-1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
