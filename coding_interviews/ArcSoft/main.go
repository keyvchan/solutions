package main

import (
	"fmt"
	"io"
)

func main() {
	var m, n int

	for {

		// read from stdin
		_, err := fmt.Scan(&m, &n)
		if err == io.EOF {
			break
		}

		A := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&A[i])
		}

		fmt.Println(Solution(m, n, A))

	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(m int, n int, A []int, current_row int, current_col int, memo map[Key]int) int {
	if current_row == m+1 {
		// done
		return 0
	}
	result := 0

	// place to the current row and column
	new_col := current_col + 1
	new_row := current_row
	if new_col > n {
		new_row++
		new_col = 1
	}
	array := serialize(A)
	if val, ok := memo[Key{array, new_row, new_col}]; ok {
		return val
	}

	// we have two choices
	result = max(result, dfs(m, n, A[:len(A)-1], new_row, new_col, memo)+A[len(A)-1]*current_row*current_col)
	result = max(result, dfs(m, n, A[1:], new_row, new_col, memo)+A[0]*current_row*current_col)

	memo[Key{array, new_row, new_col}] = result

	return result
}
func serialize(A []int) string {
	return fmt.Sprintf("%v", A)
}

type Key struct {
	A string
	x int
	y int
}

func Solution(m, n int, A []int) int {
	memo := map[Key]int{}
	// mod
	return dfs(m, n, A, 1, 1, memo) % (1e9 + 7)
}
