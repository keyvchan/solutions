package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func backtracking(start int, combo []int, n int, k int, combos *[][]int) {
	if len(combo) == k {
		new_array := make([]int, len(combo))
		copy(new_array, combo)
		*combos = append(*combos, new_array)
		return
	}

	for i := start; i < n+1; i++ {
		combo = append(combo, i)
		backtracking(i+1, combo, n, k, combos)
		combo = combo[:len(combo)-1]
	}
}

func combine(n int, k int) [][]int {
	combo := []int{}
	combos := [][]int{}
	backtracking(1, combo, n, k, &combos)

	return combos
}
