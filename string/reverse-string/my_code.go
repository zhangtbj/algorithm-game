package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}

	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	start, end := 0, len(s)-1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
