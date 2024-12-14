package main

import "fmt"

func main() {
	s := "([)]"

	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	ss := []byte(s)
	var test []byte

	for i := range ss {
		if len(test) == 0 || !isPair(ss[i], test[len(test)-1]) {
			test = append(test, ss[i])
		} else {
			if len(test) > 1 {
				test = test[:len(test)-1]
			} else {
				test = []byte{}
			}
		}
	}

	return len(test) == 0
}

func isPair(old, new byte) bool {
	if new == '(' && old == ')' {
		return true
	}
	if new == '[' && old == ']' {
		return true
	}
	if new == '{' && old == '}' {
		return true
	}
	return false
}
