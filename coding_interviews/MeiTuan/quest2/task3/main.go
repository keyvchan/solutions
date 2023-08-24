package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solution([]int{89, 38}, []int{445, 754}, 1))
}

func dfs(A []int, B []int, index int, m int, memo map[K]float64) float64 {
	if index == len(A) {
		return 0
	}
	if val, ok := memo[K{index, m}]; ok {
		return val
	}
	if m == 0 {
		// return after index
		result := 0.0
		for i := index; i < len(A); i++ {
			result += (float64(A[i]) / 100) * float64(B[i])
		}
		memo[K{index, m}] = result
		return result
	}
	result := 0.0
	result = math.Max(result, dfs(A, B, index+1, m, memo)+(float64(A[index])/100)*float64(B[index]))
	result = math.Max(result, dfs(A, B, index+1, m-1, memo)+(float64(B[index])))
	memo[K{index, m}] = result

	return result
}

type K struct {
	index int
	m     int
}

func Solution(A []int, B []int, m int) float64 {
	memo := map[K]float64{}
	// we could rewind every m
	return dfs(A, B, 0, m, memo)

}
