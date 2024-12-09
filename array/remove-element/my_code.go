package main

import (
	"fmt"
)

func main() {
	var nums = []int{2}
	val := 3
	//var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	//val := 2

	fmt.Println(removeElementTest(nums, val))
}

// 快慢指针
// 视频方法：https://www.bilibili.com/video/BV12A4y1Z7LP?spm_id_from=333.788.videopod.sections&vd_source=f881def7ea7cf10e6fa73627efe940dd
func removeElementTest(nums []int, val int) int {
	var slow int

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
