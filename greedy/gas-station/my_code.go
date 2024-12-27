package main

//import (
//	"fmt"
//)
//
//func main() {
//	var gas = []int{1, 2, 3, 4, 5}
//	var cost = []int{3, 4, 5, 1, 2}
//
//	fmt.Println(canCompleteCircuit(gas, cost))
//}
//
//func canCompleteCircuit(gas []int, cost []int) int {
//	l := len(gas)
//	var rest = make([]int, l)
//	var restSum int
//	for i := 0; i < l; i++ {
//		rest[i] = gas[i] - cost[i]
//		restSum += rest[i]
//	}
//
//	fmt.Println(rest)
//	if restSum >= 0 {
//		pos := -1
//		sum := 0
//		for i := 0; i < l; i++ {
//			fmt.Printf("i=%d, resti=%d\n", i, rest[i])
//			if pos == -1 && rest[i] < 0 {
//				continue
//			} else {
//				if pos == -1 {
//					pos = i
//				}
//				for j := i; j < l; j++ {
//					sum += rest[j]
//					if sum < 0 {
//						i = j
//						sum = 0
//						pos = -1
//					}
//					break
//				}
//			}
//		}
//		return pos
//	}
//	return -1
//}
