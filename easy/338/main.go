package main

import "fmt"

func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 0; i < n+1; i++ {

		counts := 0
		for num := i; num != 0; {
			if num%2 == 1 {
				counts++
			}
			num = num / 2
		}
		ans[i] = counts

	}

	return ans
}

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
