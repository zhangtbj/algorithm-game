package main

import (
	"fmt"
)

func main() {
	var nums = []int{2, 1, 1, 1, 1}

	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	var count int
	l := len(nums)
	cur, next := 0, 0

	for i := 0; i < l; i++ {
		next = max(next, i+nums[i])
		if i == cur {
			if cur < l-1 {
				count++
				cur = next
				if cur >= l-1 {
					return count
				}
			} else {
				return count
			}
		}
	}
	return count
}
