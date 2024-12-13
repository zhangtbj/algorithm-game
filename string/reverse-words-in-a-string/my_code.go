package main

import (
	"fmt"
)

func main() {
	s := "the sky is blue"

	fmt.Println(reverseWordsTest(s))
}

func reverseWordsTest(s string) string {
	var result []string
	var resultStr string
	var subStr string
	isWord := false
	for i := range s {
		if !isWord && s[i] == ' ' {
			continue
		}

		if !isWord && s[i] != ' ' {
			isWord = true
			subStr = string(s[i])
			if i == len(s)-1 {
				result = append(result, subStr)
			}
			continue
		}

		if isWord && s[i] != ' ' {
			subStr += string(s[i])
			if i == len(s)-1 {
				result = append(result, subStr)
			}
			continue
		}

		if isWord && s[i] == ' ' {
			isWord = false
			result = append(result, subStr)
			continue
		}
	}

	for i := len(result) - 1; i >= 0; i-- {
		resultStr += result[i] + " "
	}

	return resultStr[:len(resultStr)-1]
}
