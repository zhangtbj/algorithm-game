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

	fmt.Println(maxValue(goods, n))
}

func maxValue(goods [][]int, n int) int {
	m := len(goods)
	dp := make([][]int, m) //step1 dp[i][j], i是可取物品的范围，j是背包的范围

	//step3 init dp
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			if i == 0 {
				if j >= goods[0][0] {
					dp[0][j] = goods[0][1]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j <= n; j++ {
			tmp := 0
			if j-goods[i][0] >= 0 {
				tmp = dp[i-1][j-goods[i][0]] + goods[i][1]
			}
			dp[i][j] = max(dp[i-1][j], tmp) //step2 不取i物品，或者取i物品的最大值
		}
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[m-1][n]
}
