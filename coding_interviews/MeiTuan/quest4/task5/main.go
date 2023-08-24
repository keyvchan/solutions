package main

import (
	"fmt"
	"io"
)

func main() {

	var n, k, T int
	for {
		_, err := fmt.Scan(&n, &k, &T)
		if err == io.EOF {
			break
		}

		times := make([]int, k)
		for i := 0; i < k; i++ {
			fmt.Scan(&times[i])
		}
		events := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&events[i])
		}

	}

}

func dfs(times []int, events []int, index int, T int, toys []int) int {
	if index == len(events) {
		return 0
	}
	// we have a array to store all the toys that are collected
	switch events[index] {
	case 0:
		{
			if len(toys) == 0 {
				return dfs(times, events, index+1) + T
			} else {
				// use the min toys
				return dfs(times, events, index+1) + times[toys[0]]
			}
		}
	}

}

func Solution(times []int, events []int, T int) int {
	toys := []int{}
	dfs(times, events, 0, T, toys)
	return 0
}
