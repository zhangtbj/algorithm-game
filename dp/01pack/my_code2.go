package main

import (
	"fmt"
)

/*
6 6
2 2 3 1 5 2
2 3 1 5 4 3
*/

func main() {
	var m, n int
	fmt.Scan(&m, &n)

	var goods = make([][]int, m)

	for i := range goods {
		goods[i] = make([]int, 2)
		var s int
		fmt.Scan(&s)
		goods[i][0] = s
	}

	for i := range goods {
		var v int
		fmt.Scan(&v)
		goods[i][1] = v
	}

	//fmt.Println(goods)
	//sort.Slice(goods, func(i, j int) bool {
	//	return goods[i][0] < goods[j][0]
	//})
	fmt.Println(goods)

	fmt.Println(maxValue2(goods, n))
}

func maxValue2(goods [][]int, n int) int {
	m := len(goods)
	dp := make([]int, n+1) //step1 dp[j], j是背包的范围

	//step3 init dp
	// 0 is default

	for i := 0; i < m; i++ {
		for j := n; j >= goods[i][0]; j-- {
			dp[j] = max(dp[j], dp[j-goods[i][0]]+goods[i][1]) //step2 不取i物品，或者取i物品的最大值
			fmt.Println(dp)
			fmt.Println("---------------------")
		}
		fmt.Println()
	}

	//step5 debug
	//fmt.Println(dp)

	return dp[n]
}
