package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1234

	fmt.Println(monotoneIncreasingDigits(n))
}

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	l := len(ss)
	if l <= 1 {
		return n
	}
	flag := l
	for i := l - 1; i > 0; i-- {
		if ss[i] < ss[i-1] {
			ss[i-1] -= 1
			flag = i
		}
	}

	for i := flag; i < l; i++ {
		ss[i] = '9'
	}

	res, _ := strconv.Atoi(string(ss))

	return res
}
