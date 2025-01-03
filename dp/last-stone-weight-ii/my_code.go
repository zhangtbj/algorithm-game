package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 4, 1, 8, 1}

	fmt.Println(lastStoneWeightII(nums))
}

func lastStoneWeightII(stones []int) int {
	var sum int
	var target int
	for i := range stones {
		sum += stones[i]
	}
	target = sum / 2
	fmt.Printf("sum: %d, target: %d\n", sum, target)

	var dp []int //step1 dp[j] as 01pack
	//step3 init dp
	l := len(stones)
	dp = make([]int, target+1)

	//step4 loop
	for i := 0; i < l; i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i]) //step2 01pack formula
			//fmt.Println(dp)
		}
		//fmt.Println()
	}

	//step5 debug
	fmt.Println(dp)

	return sum - dp[target]*2
}
