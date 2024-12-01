package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSumTest(nums))
}

// Timeout
//func threeSumTest(nums []int) [][]int {
//	var result [][]int
//	var resultMap = make(map[string]bool)
//	for i := range nums {
//		var numsMap = make(map[int]int)
//		for j := range nums {
//			if i == j {
//				continue
//			}
//			another := -(nums[i] + nums[j])
//			if _, ok := numsMap[another]; ok {
//				temp := []int{nums[i], another, nums[j]}
//				for m := 0; m < len(temp)-1; m++ { //-1 2 -1
//					for n := 1; n < len(temp)-m; n++ {
//						if temp[n-1] > temp[n] {
//							temp[n-1], temp[n] = temp[n], temp[n-1]
//						}
//					}
//				}
//				testStr := strconv.Itoa(temp[0]) + strconv.Itoa(temp[1]) + strconv.Itoa(temp[2])
//				if !resultMap[testStr] {
//					result = append(result, []int{temp[0], temp[1], temp[2]})
//					resultMap[testStr] = true
//				}
//
//			} else {
//				numsMap[nums[j]] = j
//			}
//		}
//	}
//
//	return result
//}

func threeSumTest(nums []int) [][]int {
	var result [][]int
	var resultMap = make(map[string]bool)

	for m := 0; m < len(nums)-1; m++ {
		for n := 1; n < len(nums)-m; n++ {
			if nums[n-1] > nums[n] {
				nums[n-1], nums[n] = nums[n], nums[n-1]
			}
		}
	}

	//fmt.Println(nums)

	for i := range nums {
		start := 0
		end := len(nums) - 1
		for start < end {
			if start == i {
				start++
			}
			if end == i {
				end--
			}
			if start == end {
				break
			}

			if nums[start]+nums[end]+nums[i] > 0 {
				end--
			} else if nums[start]+nums[end]+nums[i] < 0 {
				start++
			} else {
				temp := []int{nums[start], nums[end], nums[i]}
				for m := 0; m < len(temp)-1; m++ { //-1 2 -1
					for n := 1; n < len(temp)-m; n++ {
						if temp[n-1] > temp[n] {
							temp[n-1], temp[n] = temp[n], temp[n-1]
						}
					}
				}
				testStr := strconv.Itoa(temp[0]) + strconv.Itoa(temp[1]) + strconv.Itoa(temp[2])
				if !resultMap[testStr] {
					result = append(result, []int{temp[0], temp[1], temp[2]})
					resultMap[testStr] = true
				}
				start++
			}
		}
	}

	return result
}
