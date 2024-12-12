package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersectionTest(nums1, nums2))
}

// The solution of map
func intersectionTest(nums1 []int, nums2 []int) []int {
	var result []int
	var numsMap = make(map[int]bool)

	for i := range nums1 {
		numsMap[nums1[i]] = false
	}

	for i := range nums2 {
		if _, ok := numsMap[nums2[i]]; ok {
			numsMap[nums2[i]] = true
		}
	}

	for v, _ := range numsMap {
		if numsMap[v] {
			result = append(result, v)
		}
	}

	return result
}
