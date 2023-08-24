package main

import "fmt"

func main() {

	fmt.Println(Solution([]int{0, 1, 1}, []int{1, 2, 3}))
	fmt.Println(Solution([]int{1, 1, 1, 9, 9, 9, 9, 7, 8}, []int{2, 0, 3, 1, 6, 5, 4, 0, 0}))
}

func Solution(A []int, B []int) int {
	// build a adjacency list
	adj := map[int]map[int]bool{}

	for i := 0; i < len(A); i++ {
		if _, ok := adj[A[i]]; !ok {
			adj[A[i]] = map[int]bool{}
		}
		if _, ok := adj[B[i]]; !ok {
			adj[B[i]] = map[int]bool{}
		}
		adj[A[i]][B[i]] = true
		adj[B[i]][A[i]] = true
	}

	costs := map[int]int{}

	for len(adj) != 1 {

		for k, v := range adj {
			if k == 0 {
				continue
			}
			if len(v) == 1 {
				for k2 := range v {
					costs[k2] += costs[k] + costs[k]/5 + 1
					delete(adj[k2], k)
				}
				delete(adj, k)
			}
		}
	}
	return 0
}
