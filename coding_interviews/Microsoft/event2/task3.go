package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution([]int{4, 2, 3, 7}, 2, 2))
	fmt.Println(Solution([]int{10, 3, 4, 7}, 2, 3))
	fmt.Println(Solution([]int{4, 2, 5, 4, 3, 5, 1, 4, 2, 7}, 3, 2))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(A []int, index int, X int, Y int, memo map[pair]int) int {
	if X < 0 {
		return 0
	}
	if index > len(A)-1 {
		// we can't finish the rehabilitation before X < 0
		return math.MaxInt32
	}
	// check if we have already calculated the cost for this index
	if val, ok := memo[pair{index, X}]; ok {
		fmt.Println("memo", index, val)
		return val
	}
	cost := math.MaxInt32
	cost = min(cost, dfs(A, index+Y, X-1, Y, memo)+A[index])
	memo[pair{index, X}] = cost
	return cost

}

type pair struct {
	x         int
	remaining int
}

func Solution(A []int, X int, Y int) int {
	// write your code in Go (Go 1.4)
	total := math.MaxInt32

	// create a memo
	memo := map[pair]int{}
	for i := 0; i < len(A); i++ {
		cost := dfs(A, i, X-1, Y, memo)
		total = min(total, cost)
	}

	return total
}
