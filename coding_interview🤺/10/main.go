package main

import "fmt"

func main() {
	fmt.Println(fib(95))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	dp := map[int]int{}
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[n] % 1000000007
}
