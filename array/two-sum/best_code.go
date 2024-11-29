package main

import "fmt"

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	var numsMap = make(map[int]int)
	for i := range nums {
		another := target - nums[i]

		if _, ok := numsMap[another]; ok {
			return []int{numsMap[another], i}
		} else {
			numsMap[nums[i]] = i
		}
	}

	return nil
}
