package main

import "fmt"

func main() {
	n := 4

	fmt.Println(fib(n))
}

func fib(n int) int {
	var dp = make([]int, n+1) // step1
	dp[0] = 0                 // step3
	if n == 0 {
		return 0
	}
	dp[1] = 1
	if n == 1 {
		return 1
	}

	for i := 2; i <= n; i++ { // step4
		dp[i] = dp[i-1] + dp[i-2] // step2
	}

	fmt.Println(dp) // step5

	return dp[n]
}
