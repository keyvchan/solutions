package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	fmt.Println(combinationSum2([]int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12},
		27))
}

func backtracking(candidates []int, current []int, current_sum int, result *[][]int, target int) {
	if current_sum == target {
		// append current to result
		new_current := make([]int, len(current))
		copy(new_current, current)
		*result = append(*result, new_current)
		return
	}
	if current_sum > target {
		return
	}
	if len(candidates) == 0 {
		return
	}

	var prev int
	for i, v := range candidates {
		if v == prev {
			continue
		}
		// copy current

		current = append(current, v)
		backtracking(candidates[i+1:], current, current_sum+v, result, target)
		current = current[:len(current)-1]
		prev = v

	}

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	backtracking(candidates, []int{}, 0, &result, target)
	return result
}
