package main

import (
	"fmt"
)

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	fmt.Println(findOrder(2, [][]int{}))
	fmt.Println(findOrder(7, [][]int{{1, 0}, {0, 3}, {0, 2}, {3, 2}, {2, 5}, {4, 5}, {5, 6}, {2, 4}}))
}

func dfs(index int, adjacencyList map[int]map[int]bool, result *[]int, visited map[int]bool, cycle map[int]bool) bool {
	if cycle[index] {
		return false
	}
	if visited[index] {
		return true
	}

	cycle[index] = true
	for k := range adjacencyList[index] {
		if !dfs(k, adjacencyList, result, visited, cycle) {
			return false
		}
	}
	delete(cycle, index)
	visited[index] = true
	*result = append(*result, index)
	return true
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// adjacency list
	adjacencyList := make(map[int]map[int]bool)
	for i := 0; i < numCourses; i++ {
		adjacencyList[i] = map[int]bool{}
	}
	for _, preq := range prerequisites {
		adjacencyList[preq[0]][preq[1]] = true
	}

	result := make([]int, 0, numCourses)
	// topological sort
	visited := map[int]bool{}
	cycle := map[int]bool{}
	for i := 0; i < numCourses; i++ {
		// dfs on every node
		if !dfs(i, adjacencyList, &result, visited, cycle) {
			return []int{}
		}
	}

	return result
}
