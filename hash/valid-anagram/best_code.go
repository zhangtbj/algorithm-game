package main

import "fmt"

func main() {
	s := "rat"
	t := "car"

	fmt.Println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	var dic = make([]int, 26)
	for i := range s {
		dic[s[i]-'a']++
	}

	for i := range t {
		dic[t[i]-'a']--
	}

	for i := range dic {
		if dic[i] != 0 {
			return false
		}
	}

	return true
}
