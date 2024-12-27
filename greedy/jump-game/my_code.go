package main

//import (
//	"fmt"
//)
//
//func main() {
//	var nums = []int{2, 3, 1, 1, 4}
//
//	fmt.Println(canJump(nums))
//}
//
//func canJump(nums []int) bool {
//	var path = make([]bool, len(nums))
//
//	path[0] = true
//	for i := range nums {
//		if path[i] {
//			for j := 1; j <= nums[i]; j++ {
//				if i+j < len(nums) {
//					path[i+j] = true
//				}
//			}
//		}
//
//		if path[len(nums)-1] == true {
//			return true
//		}
//	}
//	return false
//}
