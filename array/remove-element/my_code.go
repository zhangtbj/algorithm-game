package main

import (
	"fmt"
)

func main() {
	var nums = []int{4, 5}
	val := 4
	//var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	//val := 2

	fmt.Println(removeElementTest(nums, val))
}

func removeElementTest(nums []int, val int) int {
	var pos, checker int

	for pos <= len(nums)-1 {
		for nums[checker] == val {
			checker++
			if checker == len(nums) {
				return pos
			}
		}
		nums[pos], nums[checker] = nums[checker], nums[pos]
		pos++
		checker++

		if checker >= len(nums) {
			break
		}
	}

	return pos
}
