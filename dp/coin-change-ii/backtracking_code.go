package main

import (
	"fmt"
	"sort"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 5

	fmt.Println(change2(amount, coins))
}

var count int

func change2(amount int, coins []int) int {
	count = 0
	sort.Ints(coins) // 排序，为剪枝做准备
	backtracking(coins, amount, 0, 0)
	return count
}

func backtracking(coins []int, amount int, sum int, startIndex int) {
	if amount == sum {
		count++
		//fmt.Printf("count: %d\n", count)
		return
	}

	if amount < sum {
		return
	}

	l := len(coins)

	for i := startIndex; i < l && sum < amount; i++ {
		sum += coins[i]
		//fmt.Printf("coins[i]: %d, sum: %d\n", coins[i], sum)
		backtracking(coins, amount, sum, i)
		sum -= coins[i]
		fmt.Println()
	}
}
