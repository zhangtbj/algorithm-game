package main

import (
	"fmt"
)

func main() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m := 5
	n := 3

	fmt.Println(findMaxForm(strs, m, n))
}

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	var strs0 = make([]int, l)
	var strs1 = make([]int, l)
	for i := range strs {
		strs0[i] = getCount(strs[i], '0')
		strs1[i] = getCount(strs[i], '1')
	}

	// fmt.Println(strs0)
	// fmt.Println(strs1)

	//step1 dp为有m个0，n个1的个数
	var dp = make([][]int, m+1)

	//step3 default is 0
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	//step4 loop
	for x := 0; x < l; x++ {
		for i := m; i >= strs0[x]; i-- {
			for j := n; j >= strs1[x]; j-- {
				//fmt.Printf("i: %d, j: %d, strs0[i]: %d, strs[1]: %d\n", i, j, strs0[i], strs1[i])
				dp[i][j] = max(dp[i][j], dp[i-strs0[x]][j-strs1[x]]+1) //step2 formula
			}
		}
	}

	//step5 debug
	for i := range dp {
		fmt.Println(dp[i])
	}

	return dp[m][n]
}

func getCount(str string, target byte) int {
	var count int
	for i := range str {
		if str[i] == target {
			count++
		}
	}
	return count
}
