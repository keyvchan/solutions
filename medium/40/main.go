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

func backtracking(candidates []int, current []int, current_sum int, last int, result *[][]int, target int) {
	fmt.Println(candidates)
	if current_sum == target {
		// append current to result
		if len(candidates) == 0 {
			*result = append(*result, current)
			return
		}
		if candidates[0] != last {
			*result = append(*result, current)
			return
		}
		return
	}
	if current_sum > target {
		return
	}
	if len(candidates) == 0 {
		return
	}

	// copy current
	new_current := make([]int, len(current))
	copy(new_current, current)
	backtracking(candidates[1:], append(new_current, candidates[0]), current_sum+candidates[0], candidates[0], result, target)

	if candidates[0] != last || last == 0 {
		backtracking(candidates[1:], current, current_sum, last, result, target)
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	backtracking(candidates, []int{}, 0, 0, &result, target)
	return result
}
