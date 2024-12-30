package main

import "fmt"

func main() {
	//nums := []int{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 97}
	nums := []int{100, 100, 99, 97}

	fmt.Println(canPartition2(nums))
}

func canPartition2(nums []int) bool {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	} else {
		sum /= 2
	}

	fmt.Printf("sum: %d\n", sum)

	return backtracking(nums, 0, 0, sum)
}

// report timeout for backtracking method
func backtracking(nums []int, startIndex int, sum int, check int) bool {
	if sum > check {
		return false
	}
	if sum == check {
		return true
	}

	for i := startIndex; i < len(nums); i++ {
		sum += nums[i]
		fmt.Printf("sum in: %d\n", sum)
		if backtracking(nums, i+1, sum, check) {
			return true
		}
		sum -= nums[i]
	}
	return false
}
