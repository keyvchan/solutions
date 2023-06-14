package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	last := 2
	second := 1
	current := 0

	// count n-1, store n-2
	for i := 3; i < n; i++ {
		current = last + second
		second = last
		last = current

	}

	return last + second
}
