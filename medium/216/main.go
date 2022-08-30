package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(4, 1))
}

func backtracking(k int, n int, current_i int, current []int, result *[][]int, used map[int]bool) {
	if len(current) == k {
		// we can check if its sums up to n
		if n == 0 {
			*result = append(*result, current)
		}
		return
	}

	for i := current_i; i <= 9; i++ {
		if used[i] {
			continue
		}
		new_used := map[int]bool{}
		for key, value := range used {
			new_used[key] = value
		}
		new_used[i] = true
		new_current := make([]int, len(current))
		copy(new_current, current)
		new_current = append(new_current, i)
		backtracking(k, n-i, i, new_current, result, new_used)
	}

}

func combinationSum3(k int, n int) [][]int {

	used := map[int]bool{}

	result := [][]int{}

	backtracking(k, n, 1, []int{}, &result, used)

	return result
}
