package main

import (
	"fmt"
)

func main() {
	var gas = []int{1, 2, 3, 4, 5}
	var cost = []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	var rest = make([]int, l)
	var restSum int
	var pos = 0
	var sum = 0
	for i := 0; i < l; i++ {
		rest[i] = gas[i] - cost[i]
		restSum += rest[i]
	}

	fmt.Println(rest)
	if restSum >= 0 {
		for i := 0; i < l; i++ {
			sum += rest[i]
			if sum < 0 {
				pos = i + 1
				sum = 0
			}
		}
		return pos
	}
	return -1
}
