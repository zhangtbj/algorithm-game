package main

import (
	"fmt"
	"sort"
)

func main() {
	var s = []int{1, 1}
	var g = []int{1, 2, 3}

	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var sum int

	startIndex := 0
	for i := 0; i < len(s); i++ {
		for j := startIndex; j < len(g); j++ {
			//fmt.Printf("s1: %d, gj: %d\n", i, j)
			if s[i] >= g[j] {
				sum++
				startIndex++
				break
			}
		}
	}

	return sum
}
