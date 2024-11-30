package main

import "fmt"

func main() {
	var nums1 = []int{1, 3, 5}
	var nums2 = []int{2, 4, 6}

	fmt.Println(findMedianSortedArraysTest2(nums1, nums2))
}

func findMedianSortedArraysTest2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArraysTest2(nums2, nums1)
	}

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0

	for low <= high {
		nums1Mid = low + (high-low)>>1
		nums2Mid = k - nums1Mid

		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums2[nums2Mid-1] > nums1[nums1Mid] {
			low = nums1Mid + 1
		} else {
			break
		}
	}

	//fmt.Println(nums1Mid)
	//fmt.Println(nums2Mid)

	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}

	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}

	return float64(midLeft+midRight) / 2
}
