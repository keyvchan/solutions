package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}

func backtracking(candidates []int, target int, sum int, current []int, result *[][]int) {
	if sum == target {
		*result = append(*result, current)
		return
	} else if sum > target {
		return
	}
	if len(candidates) == 0 {
		return
	}

	// choose one
	// copy current
	new_current := make([]int, len(current))
	copy(new_current, current)
	new_current = append(new_current, candidates[0])
	backtracking(candidates, target, sum+candidates[0], new_current, result)
	// not choose one
	backtracking(candidates[1:], target, sum, current, result)
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}

	backtracking(candidates, target, 0, []int{}, &result)

	return result

}
