package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	couns_map := map[int]int{}
	// count n-1, store n-2
	couns_map[1] = 1
	couns_map[2] = 2
	for i := 3; i < n; i++ {
		couns_map[i] = couns_map[i-1] + couns_map[i-2]
	}

	return couns_map[n-1] + couns_map[n-2]
}
