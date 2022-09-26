package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
}

func dfs(n int, memo map[int]int) int {
	if n <= 1 {
		return 1
	}
	if v, ok := memo[n]; ok {
		return v
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += dfs(i-1, memo) * dfs(n-i, memo)
	}
	memo[n] = sum
	return sum
}

func numTrees(n int) int {
	memo := map[int]int{}

	return dfs(n, memo)

}
