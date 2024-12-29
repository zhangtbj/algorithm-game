package main

import (
	"fmt"
)

func main() {
	var ratings = []int{1, 0, 2}

	fmt.Println(candy(ratings))
}

func candy(ratings []int) int {
	l := len(ratings)
	isCandy := make([]int, l)
	candySum := l

	if l == 1 {
		return candySum
	}

	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			isCandy[i] = isCandy[i-1] + 1
		}
	}

	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			isCandy[i] = max(isCandy[i], isCandy[i+1]+1)
		}
	}

	fmt.Println(isCandy)
	fmt.Println(candySum)
	for i := range isCandy {
		candySum += isCandy[i]
	}
	fmt.Println(candySum)
	return candySum
}
