package main

import "fmt"

func main() {
	var nums = []int{-4, -1, 0, 3, 10}

	fmt.Println(sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	var result = make([]int, len(nums))
	k := len(nums) - 1
	for i, j := 0, len(nums)-1; i <= j; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			i++
		} else {
			result[k] = nums[j] * nums[j]
			j--
		}
	}

	return result
}
