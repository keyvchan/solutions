package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	// the newInterval have three situations
	// it could be before the current interval
	// it could be after the current interval
	// it could be inside the current interval
	result := make([][]int, 0, len(intervals)+1)

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			// check new interval is before the current interval
			fmt.Println("newInterval[1] < intervals[i][0]", newInterval[1], intervals[i][0])
			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		} else if newInterval[0] > intervals[i][1] {
			// check newInterval is after the current interval
			// insert the current interval into the result, and continue
			result = append(result, intervals[i])
		} else {
			// check newInterval is inside the current interval
			// newInterval[1] > intervals[i][0]
			// newInterval[0] < intervals[i][1]
			if newInterval[0] > intervals[i][0] {
				newInterval[0] = intervals[i][0]
			}

			if newInterval[1] < intervals[i][1] {
				newInterval[1] = intervals[i][1]
			}
		}

	}
	result = append(result, newInterval)

	return result
}
