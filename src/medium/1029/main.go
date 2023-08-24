package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
	fmt.Println(twoCitySchedCost([][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}))
	fmt.Println(twoCitySchedCost([][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(costs [][]int, ith int, A int, B int, memo map[pair]int) int {
	if A > len(costs)/2 || B > len(costs)/2 {
		// we have more than n people in A or B, the path is invalid
		return math.MaxInt32
	}
	if ith == len(costs) {
		return 0
	}
	if _, ok := memo[pair{A, B}]; ok {
		return memo[pair{A, B}]
	}

	// we could choose A or B
	total_cost := math.MaxInt32
	total_cost = min(total_cost, costs[ith][0]+dfs(costs, ith+1, A+1, B, memo))
	total_cost = min(total_cost, costs[ith][1]+dfs(costs, ith+1, A, B+1, memo))

	memo[pair{A, B}] = total_cost
	return total_cost
}

type pair struct {
	A int
	B int
}

func twoCitySchedCost(costs [][]int) int {
	memo := map[pair]int{}

	return dfs(costs, 0, 0, 0, memo)

}
