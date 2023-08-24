package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(satisfaction []int, time int, index int, memo map[Key]int) int {

	if index == len(satisfaction) {
		return 0
	}

	maxLikeTimeCoefficient := 0

	if maxLikeTimeCoefficient, ok := memo[Key{Index: index, Time: time}]; ok {
		return maxLikeTimeCoefficient
	}
	// eat this dish
	maxLikeTimeCoefficient = max(maxLikeTimeCoefficient, dfs(satisfaction, time+1, index+1, memo)+satisfaction[index]*time)
	// or not eat this dish
	maxLikeTimeCoefficient = max(maxLikeTimeCoefficient, dfs(satisfaction, time, index+1, memo))

	memo[Key{Index: index, Time: time}] = maxLikeTimeCoefficient

	return maxLikeTimeCoefficient
}

type Key struct {
	Index int
	Time  int
}

func maxSatisfaction(satisfaction []int) int {

	// fist we sort the satisfaction array
	sort.Ints(satisfaction)

	memo := map[Key]int{}

	return dfs(satisfaction, 1, 0, memo)
}
