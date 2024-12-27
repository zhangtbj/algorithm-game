package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 7, 4, 9, 2, 5}

	fmt.Println(wiggleMaxLength(nums))
}

func wiggleMaxLength(nums []int) int {
	var result int

	l := len(nums)
	if l < 2 {
		return l
	}

	var larger int
	for i := 1; i < l; i++ {
		diffNum := nums[i] - nums[i-1]
		if (larger == 0 && diffNum != 0) || ((diffNum > 0 && larger == 2) || (diffNum < 0 && larger == 1)) {
			result++
		}

		if diffNum > 0 {
			larger = 1
		} else if diffNum < 0 {
			larger = 2
		}
	}

	return result + 1
}
