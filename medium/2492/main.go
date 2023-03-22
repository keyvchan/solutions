package main

import (
	"sort"
)

func main() {

}

func dfs(node int, roadsMap map[int][]int, visited map[int]bool) {
	if visited[node] {
		return
	}

	visited[node] = true

	for _, next := range roadsMap[node] {
		dfs(next, roadsMap, visited)
	}

}

func minScore(_ int, roads [][]int) int {

	roadsMap := map[int][]int{}

	for _, road := range roads {
		roadsMap[road[0]] = append(roadsMap[road[0]], road[1])
		roadsMap[road[1]] = append(roadsMap[road[1]], road[0])
	}

	// start from 1, let's see how many node we can reach
	visited := map[int]bool{}

	dfs(1, roadsMap, visited)

	sort.Slice(roads, func(i, j int) bool {
		return roads[i][2] < roads[j][2]
	})

	for _, edge := range roads {

		if visited[edge[0]] || visited[edge[1]] {
			return edge[2]
		}
	}

	return 0
}
