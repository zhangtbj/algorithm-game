package main

import (
	"fmt"
)

func main() {
	//var nums = []int{4, 5}
	//val := 4
	var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {
	pos := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if pos != i {
				nums[pos], nums[i] = nums[i], nums[pos]
			}
			pos++
		}
	}

	return pos
}
