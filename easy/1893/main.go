package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isCovered([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, 7, 9))
	fmt.Println(isCovered([][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5))
	fmt.Println(isCovered([][]int{{1, 10}, {10, 20}}, 21, 21))
	fmt.Println(isCovered([][]int{{1, 50}}, 1, 50))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isCovered(ranges [][]int, left int, right int) bool {
	// sort by left
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	newRanges := [][]int{}
	// merge all ranges
	left_range := ranges[0][0]
	right_range := ranges[0][1]
	for i, r := range ranges {
		// since we sort it, all ranges must be after the current one
		if r[0] > right_range+1 {
			// append the current range
			newRanges = append(newRanges, []int{left_range, right_range})
			// after
			left_range = r[0]
			right_range = r[1]
		} else {
			// inside the current one, extend current one
			right_range = max(r[1], right_range)
		}
		if i == len(ranges)-1 {
			newRanges = append(newRanges, []int{left_range, right_range})
		}
	}
	fmt.Println(newRanges)

	for _, r := range newRanges {
		// check if contains the left and right
		if r[0] <= left && r[1] >= right {
			return true
		}
	}

	return false

}
