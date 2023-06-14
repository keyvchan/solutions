package main

import "fmt"

func main() {
	fmt.Println(numDecodings("1201234"))
	fmt.Println(numDecodings("111111111111111111111111111111111111111111111"))
}

func numDecodings(s string) int {
	dp := map[int]int{}
	dp[len(s)] = 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}
		if i < len(s)-1 {
			if s[i] == '1' || (s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6') {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}
