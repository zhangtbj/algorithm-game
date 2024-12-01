package main

import "fmt"

func main() {
	var nums = []int{2, 3, 1, 3, 3}

	fmt.Println(maxAreaTest(nums))
}

// Timeout
//func maxAreaTest(height []int) int {
//	if len(height) == 2 {
//		if height[0] > height[1] {
//			return height[1]
//		} else {
//			return height[0]
//		}
//	} else {
//		right := height[:len(height)-1]
//		left := height[1:]
//		var maxH int
//		if height[0] > height[len(height)-1] {
//			maxH = height[len(height)-1] * (len(height) - 1)
//		} else {
//			maxH = height[0] * (len(height) - 1)
//		}
//		return max(maxH, maxAreaTest(right), maxAreaTest(left))
//	}
//}

// Out of memory
func maxAreaTest(height []int) int {
	lenH := len(height)
	var matrix = make([][]int, lenH)
	for i := range matrix {
		matrix[i] = make([]int, lenH)
	}
	var maxS int

	for i := 1; i < lenH; i++ {
		for j := 0; j < lenH; j++ {
			var s1, s2, maxH int
			if j-i >= 0 {
				if height[j] > height[j-i] {
					maxH = height[j-i]
				} else {
					maxH = height[j]
				}
				s1 = max(matrix[i-1][j], maxH*i)
			}
			if j+i < lenH {
				if height[j] > height[j+i] {
					maxH = height[j+i]
				} else {
					maxH = height[j]
				}
				s2 = max(matrix[i-1][j], maxH*i)
			}
			middleS := max(matrix[i-1][j], s1, s2)
			matrix[i][j] = middleS
			if maxS < middleS {
				maxS = middleS
			}
		}
	}

	//fmt.Println(matrix)
	return maxS
}
