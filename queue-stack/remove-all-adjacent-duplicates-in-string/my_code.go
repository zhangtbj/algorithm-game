package main

import "fmt"

func main() {
	s := "abbaca"

	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	ss := []byte(s)
	var test []byte
	for i := range ss {
		if len(test) == 0 || ss[i] != test[len(test)-1] {
			test = append(test, ss[i])
		} else {
			test = test[:len(test)-1]
		}
	}

	return string(test)
}
