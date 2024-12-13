package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl"
	k := 39

	fmt.Println(reverseStrTest(s, k))
}

func reverseStrTest(s string, k int) string {
	strArr := strings.Split(s, "")
	var news string
	var start, end int
	for i := 0; i < len(s); i += 2 * k {
		if len(s)-i >= k {
			start, end = i, i+k-1
		} else {
			start, end = i, len(s)-1
		}
		for start < end {
			strArr[start], strArr[end] = strArr[end], strArr[start]
			start++
			end--
		}
	}

	for i := range strArr {
		news += strArr[i]
	}

	return news
}
