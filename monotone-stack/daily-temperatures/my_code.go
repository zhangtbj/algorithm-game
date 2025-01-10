package main

import (
	"fmt"
)

func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	fmt.Println(dailyTemperatures(temperatures))
}

func dailyTemperatures(temperatures []int) []int {
	l := len(temperatures)
	var result = make([]int, l)

	var stack []int
	stack = append(stack, 0)

	for i := 1; i < l; i++ {
		if len(stack) > 0 {
			//fmt.Println(temperatures[i])
			if temperatures[i] > temperatures[stack[len(stack)-1]] {
				for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
					//fmt.Println(stack)
					result[stack[len(stack)-1]] = i - stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
			}
			stack = append(stack, i)
		}
	}

	return result
}
