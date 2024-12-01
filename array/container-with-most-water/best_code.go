package main

import "fmt"

func main() {
	var nums = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(nums))
}

func maxArea(height []int) int {
	low := 0
	high := len(height) - 1
	var maxS int

	for low <= high {
		var temp int
		if height[low] < height[high] {
			temp = height[low] * (high - low)
			low++
		} else {
			temp = height[high] * (high - low)
			high--
		}

		if maxS < temp {
			maxS = temp
		}
	}

	return maxS
}
