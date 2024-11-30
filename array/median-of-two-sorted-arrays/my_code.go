package main

import "fmt"

func main() {
	var nums1 = []int{1, 3, 5}
	var nums2 = []int{2, 4, 6}

	fmt.Println(findMedianSortedArraysTest(nums1, nums2))
}

func findMedianSortedArraysTest(nums1 []int, nums2 []int) float64 {
	var twoMedian bool
	var medianPos int
	len1 := len(nums1)
	len2 := len(nums2)
	if (len1+len2)%2 == 0 {
		twoMedian = true
	}
	medianPos = (len1 + len2) / 2

	var nums3 []int

	var p1, p2 int
	for len(nums3) < medianPos+1 {

		if p1 >= len1 {
			nums3 = append(nums3, nums2[p2])
			p2++
			continue
		}
		if p2 >= len2 {
			nums3 = append(nums3, nums1[p1])
			p1++
			continue
		}

		if nums1[p1] < nums2[p2] {
			nums3 = append(nums3, nums1[p1])
			p1++
		} else {
			nums3 = append(nums3, nums2[p2])
			p2++
		}
	}

	if twoMedian {
		return (float64)(nums3[len(nums3)-1]+nums3[len(nums3)-2]) / 2
	} else {
		return (float64)(nums3[len(nums3)-1])
	}
}
