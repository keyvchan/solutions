package main

import "fmt"

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		money := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&money[i])
		}
		fmt.Print(solution(money))

	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(money []int, root int, memo map[int]int) int {
	if root >= len(money)+1 {
		return 0
	}
	if _, ok := memo[root]; ok {
		return memo[root]
	}
	// we could choose left and right child
	current_money := 0
	current_money = max(current_money, dfs(money, root*2, memo)+money[root-1])
	current_money = max(current_money, dfs(money, root*2+1, memo)+money[root-1])
	memo[root] = current_money
	return current_money
}

func solution(money []int) int {
	memo := map[int]int{}
	return dfs(money, 1, memo)

}
