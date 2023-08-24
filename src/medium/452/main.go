package main

import (
	"fmt"
	"sort"
)

func main() {
	// points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	// points := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(findMinArrowShots(points))
}

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	right := points[0][1]
	count := 0

	for i, v := range points {
		if i == 0 {
			right = v[1]
			count++
			continue
		}

		if v[0] > right {
			right = v[1]
			count++
		}
	}

	return count
}
