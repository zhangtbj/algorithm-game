package main

import "fmt"

func main() {
	s := "rat"
	t := "car"

	fmt.Println(isAnagramTest(s, t))
}

func isAnagramTest(s string, t string) bool {
	var dic = make(map[byte]int)
	for i := range s {
		dic[s[i]]++
	}

	for i := range t {
		if dic[t[i]] > 0 {
			dic[t[i]]--
		} else {
			return false
		}
	}

	for v, _ := range dic {
		if v > 0 {
			return false
		}
	}

	return true
}
