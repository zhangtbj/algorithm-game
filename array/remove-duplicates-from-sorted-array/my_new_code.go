package main

import (
	"fmt"
)

func main() {
	var nums = []int{1, 2, 3, 4, 4}

	fmt.Println(removeDuplicatesTest2(nums))
}

func removeDuplicatesTest2(nums []int) int {
	var count, checker int

	//for i := 0; i < len(nums); i++ {
	for count < len(nums)-1 {
		//for j := 0; j < len(nums); j++ {
		for nums[count] == nums[checker] {
			//if nums[count] == nums[checker] {
			checker++
			if checker == len(nums) {
				return count + 1
			}
			//continue
			//}
			count++
			nums[count] = nums[checker]
		}
	}

	return count
}
