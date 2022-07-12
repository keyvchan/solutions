package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
}

func merge(intervals [][]int) [][]int {

	// sort the intervals by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	newIntervals := [][]int{}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	currentInterval := intervals[0]
	for _, interval := range intervals {
		if len(newIntervals) == 0 {
			// add the first interval
			newIntervals = append(newIntervals, interval)
		} else {
			// interval after the last interval
			if interval[0] > currentInterval[1] {
				// add the interval
				newIntervals = append(newIntervals, interval)
				currentInterval = interval
			} else {
				// merge the intervals
				currentInterval[1] = max(currentInterval[1], interval[1])
			}
		}

	}
	return newIntervals

}
