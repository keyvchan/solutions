package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(3, [][]int{{1, 0}, {2, 1}, {3, 2}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
}

func dfs(nibers []int, adjacencyList map[int][]int, checked map[int]bool, can_finish map[int]bool) bool {
	if len(nibers) == 0 {
		return true
	}

	result := true
	for _, niber := range nibers {
		if _, ok := checked[niber]; ok {
			return false
		}
		if _, ok := can_finish[niber]; ok {
			return true
		}
		// check if nibers can be finished
		checked[niber] = true
		result = result || dfs(adjacencyList[niber], adjacencyList, checked, can_finish)
		delete(checked, niber)
	}

	return result
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// construct adjacency list
	adjacencyList := map[int][]int{}
	for _, prereq := range prerequisites {
		adjacencyList[prereq[0]] = append(adjacencyList[prereq[0]], prereq[1])
	}
	fmt.Println(adjacencyList)
	result := false
	checked := map[int]bool{}
	can_finish := map[int]bool{}
	for i := 0; i < numCourses; i++ {
		checked[i] = true
		result = result || dfs(adjacencyList[i], adjacencyList, checked, can_finish)
		delete(checked, i)
	}

	return result
}
