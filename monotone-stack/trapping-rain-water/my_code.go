package main

import (
	"fmt"
)

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(height))
}

func trap(height []int) int {
	l := len(height)
	var result = make([]int, l)
	var stack []int
	var sum int
	stack = append(stack, 0)

	for i := 1; i < l; i++ {
		if height[i] >= height[stack[len(stack)-1]] {
			for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
				mid := height[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					h := min(height[i], height[stack[len(stack)-1]]) - mid
					w := i - stack[len(stack)-1] - 1
					result[stack[len(stack)-1]] = h * w
					sum += h * w
				}
			}
		}
		stack = append(stack, i)
	}

	fmt.Println(result)

	return sum
}
