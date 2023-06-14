package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15}))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(days []int, costs []int, currentIndex int, currentPassEnd int, memo map[memoKey]int) int {
	// at days 1, we can have
	if currentIndex == len(days) {
		return 0
	}

	// if we have a memoized result
	if val, ok := memo[memoKey{currentIndex, currentPassEnd}]; ok {
		return val
	}

	// we need a new pass
	cost := math.MaxInt32

	// if we have a pass that covers this day
	if days[currentIndex] <= currentPassEnd {
		cost = dfs(days, costs, currentIndex+1, currentPassEnd, memo)
	} else {
		// buy 1 day pass
		cost = min(cost, dfs(days, costs, currentIndex+1, currentPassEnd, memo)+costs[0])
		// buy 7 day pass
		cost = min(cost, dfs(days, costs, currentIndex+1, days[currentIndex]+6, memo)+costs[1])
		// buy 30 day pass
		cost = min(cost, dfs(days, costs, currentIndex+1, days[currentIndex]+29, memo)+costs[2])
	}

	memo[memoKey{currentIndex, currentPassEnd}] = cost
	return cost

}

type memoKey struct {
	currentIndex   int
	currentPassEnd int
}

func mincostTickets(days []int, costs []int) int {
	// try every options
	memo := map[memoKey]int{}

	return dfs(days, costs, 0, 0, memo)

}
