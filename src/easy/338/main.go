package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i < n+1; i++ {

		counts := 0

		ans[i] = counts + ans[i/2] + i%2

	}

	return ans
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
