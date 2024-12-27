package main

//import (
//	"fmt"
//)
//
//func main() {
//	var nums = []int{4, -1, -1, 4, 2, -3, 4, -10, -1, 5}
//
//	fmt.Println(maxSubArray(nums))
//}
//
//func maxSubArray(nums []int) int {
//	var result = -10005
//
//	if len(nums) == 1 {
//		return nums[0]
//	}
//
//	var newNums []int
//	var larger = false
//	var item = nums[0]
//	var always = true
//	var alwaysSum = nums[0]
//	var largerItem = nums[0]
//	if nums[0] >= 0 {
//		larger = true
//	}
//
//	for i := 1; i < len(nums); i++ {
//		if (nums[i] >= 0 && larger) || (nums[i] < 0 && !larger) {
//			item += nums[i]
//			alwaysSum += nums[i]
//			largerItem = max(largerItem, nums[i])
//		} else {
//			newNums = append(newNums, item)
//			item = nums[i]
//			if nums[i] >= 0 {
//				larger = true
//			} else {
//				larger = false
//			}
//			always = false
//		}
//		if i == len(nums)-1 {
//			newNums = append(newNums, item)
//		}
//	}
//	fmt.Println(newNums)
//
//	if always {
//		if nums[0] >= 0 {
//			return alwaysSum
//		} else {
//			return largerItem
//		}
//	}
//
//	for i := 0; i < len(newNums); i++ {
//		maxSum := newNums[i]
//		sum := newNums[i]
//		for j := i + 1; j < len(newNums); j++ {
//			sum += newNums[j]
//			//fmt.Printf("sum: %d, newNums: %d\n", sum, newNums[j])
//			maxSum = max(maxSum, sum)
//		}
//
//		result = max(result, maxSum)
//		//fmt.Printf("result: %d\n\n", result)
//	}
//
//	return result
//}
