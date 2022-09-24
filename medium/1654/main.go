package main

import "fmt"

func main() {
	fmt.Println(minimumJumps([]int{14, 4, 18, 1, 15}, 3, 15, 9))
	fmt.Println(minimumJumps([]int{8, 3, 16, 6, 12, 20}, 15, 13, 11))
	fmt.Println(minimumJumps([]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7))
	fmt.Println(minimumJumps([]int{
		162, 118, 178, 152, 167, 100, 40, 74,
		199, 186, 26, 73, 200, 127, 30,
		124, 193, 84, 184, 36, 103, 149,
		153, 9, 54, 154, 133, 95, 45, 198,
		79, 157, 64, 122, 59, 71, 48, 177,
		82, 35, 14, 176, 16, 108, 111, 6, 168, 31,
		134, 164, 136, 72, 98,
	},
		29,
		98,
		80,
	))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(forbidden map[int]bool, current_index int, a, b, x int, forward bool, max_forbidden int, memo map[pair]int, visited map[int]bool) int {
	if forbidden[current_index] {
		// invalid jump
		return -1
	}
	if current_index < 0 {
		// invalid jump
		return -1
	}
	if visited[current_index] {
		return -1
	}
	if current_index > a+b+max(x, max_forbidden) {
		// invalid jump
		return -1
	}
	if current_index == x {
		return 0
	}
	if val, ok := memo[pair{current_index, forward}]; ok {
		return val
	}
	first := -1
	second := -1

	if !forward {
		visited[current_index] = true
		first = dfs(forbidden, current_index+a, a, b, x, true, max_forbidden, memo, visited)
		delete(visited, current_index)
	} else {
		visited[current_index] = true
		first = dfs(forbidden, current_index+a, a, b, x, true, max_forbidden, memo, visited)
		delete(visited, current_index)
		visited[current_index] = true
		second = dfs(forbidden, current_index-b, a, b, x, false, max_forbidden, memo, visited)
		delete(visited, current_index)
	}
	if first == -1 && second == -1 {
		memo[pair{current_index, forward}] = -1
		return -1
	}
	if first == -1 {
		memo[pair{current_index, forward}] = second + 1
		return second + 1
	}
	if second == -1 {
		memo[pair{current_index, forward}] = first + 1
		return first + 1
	}
	minium := min(first, second) + 1
	memo[pair{current_index, forward}] = minium
	return minium

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pair struct {
	index   int
	forward bool
}

func minimumJumps(forbidden []int, a int, b int, x int) int {
	memo := map[pair]int{}
	max_forbidden := 0
	// create a map of forbidden indices for O(1) lookup
	forbidden_map := make(map[int]bool)
	for _, index := range forbidden {
		if index > max_forbidden {
			max_forbidden = index
		}
		forbidden_map[index] = true
	}

	visited := map[int]bool{}

	return min(dfs(forbidden_map, 0, a, b, x, true, max_forbidden, memo, visited), dfs(forbidden_map, 0, a, b, x, false, max_forbidden, memo, visited))
}
