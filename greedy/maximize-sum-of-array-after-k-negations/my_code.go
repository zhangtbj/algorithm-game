package main

//import (
//	"fmt"
//	"sort"
//)
//
//func main() {
//	var nums = []int{3, -1, 0, 2}
//
//	k := 3
//
//	fmt.Println(largestSumAfterKNegations(nums, k))
//}
//
//func largestSumAfterKNegations(nums []int, k int) int {
//	var sum int
//	l := len(nums)
//	sort.Ints(nums)
//
//	for i := 0; i < l; i++ {
//		sum += nums[i]
//	}
//
//	if k < l && nums[k-1] <= 0 {
//		for i := 0; i < k; i++ {
//			if i < k {
//				sum -= 2 * nums[i]
//			}
//		}
//	} else {
//		if nums[0] <= 0 {
//			for i := 0; i < k; i++ {
//				if i < l && nums[i] < 0 {
//					sum -= 2 * nums[i]
//				} else {
//					if (k-i)%2 != 0 {
//						if i < l && nums[i]+nums[i-1] < 0 {
//							sum -= 2 * nums[i]
//						} else {
//							sum += 2 * nums[i-1]
//						}
//					}
//					break
//				}
//			}
//		} else {
//			if k%2 != 0 {
//				sum -= 2 * nums[0]
//			}
//		}
//	}
//	return sum
//}
