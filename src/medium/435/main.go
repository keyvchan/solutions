package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println(eraseOverlapIntervals([][]int{
		{-52, 31},
		{-73, -26},
		{82, 97},
		{-65, -11},
		{-62, -49},
		{95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	}))
}

func eraseOverlapIntervals(intervals [][]int) int {

	// find max number of intervals that can possiblly be not overlapped
	// total minus the number of intervals tha can be not overlapped
	sort.Slice(intervals, func(i, j int) bool {

		return intervals[i][1] < intervals[j][1]

	})

	fmt.Println(intervals)

	right := intervals[0][1]

	count := 0

	for i, v := range intervals {
		if i == 0 {
			fmt.Println(i, v)
			right = v[1]
			count++
			continue
		}
		if v[1] > right && v[0] >= right {
			right = v[1]
			fmt.Println(i, v)
			count++
		}
	}

	return len(intervals) - count
}
