package main

import (
	"fmt"
)

func main() {
	var bills = []int{5, 5, 5, 10, 20}

	fmt.Println(lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	l := len(bills)
	var billsList = make([]int, 4)

	for i := 0; i < l; i++ {
		bill := bills[i] / 10
		if bill == 0 {
			billsList[0] += 1
		} else if bill == 1 {
			if billsList[0] > 0 {
				billsList[1] += 1
				billsList[0] -= 1
			} else {
				return false
			}
		} else {
			if billsList[0] > 0 && billsList[1] > 0 {
				billsList[2] += 1
				billsList[1] -= 1
				billsList[0] -= 1
			} else if billsList[0] >= 3 && billsList[1] == 0 {
				billsList[2] += 1
				billsList[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}
