package main

import "fmt"

func main() {
	var nums = []int{-4, -1, 0, 3, 10}

	fmt.Println(sortedSquaresTest(nums))
}

func sortedSquaresTest(nums []int) []int {
	minVal, minPos := 100000000, 0
	var result []int
	for i := range nums {
		if nums[i]*nums[i] < minVal {
			minVal = nums[i] * nums[i]
			minPos = i
		}
	}

	left, right := minPos-1, minPos+1
	result = append(result, nums[minPos]*nums[minPos])
	for left >= 0 || right < len(nums) {
		if left >= 0 && ((right < len(nums) && nums[left]*nums[left] <= nums[right]*nums[right]) || right >= len(nums)) {
			result = append(result, nums[left]*nums[left])
			left--
		} else if right < len(nums) && ((left >= 0 && nums[left]*nums[left] > nums[right]*nums[right]) || left < 0) {
			result = append(result, nums[right]*nums[right])
			right++
		}
	}

	return result
}
