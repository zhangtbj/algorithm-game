package main

import (
	"fmt"
)

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}

	fmt.Println(largestRectangleArea(heights))
}

func largestRectangleArea(heights []int) int {
	newHeights := []int{0}
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, 0)

	l := len(newHeights)
	var maxS int
	var stack []int
	stack = append(stack, 0)
	var result = make([]int, l)

	for i := 1; i < l; i++ {
		if newHeights[i] < newHeights[stack[len(stack)-1]] {
			for len(stack) > 0 && newHeights[i] < newHeights[stack[len(stack)-1]] {
				mid := newHeights[stack[len(stack)-1]]
				result[stack[len(stack)-1]] = i
				stack = stack[:len(stack)-1]
				w := i - stack[len(stack)-1] - 1
				maxS = max(maxS, w*mid)
			}
		}
		stack = append(stack, i)
	}

	fmt.Println(result)

	return maxS
}
