package main

import (
	"fmt"
	"io"
)

func main() {
	var n, b int

	for {
		_, err := fmt.Scan(&n, &b)
		if err == io.EOF {
			break
		}
		battery := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&battery[i])
		}

		time := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&time[i])
		}
		fmt.Println(Solution(battery, time, b))
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(time []int, index int, current_time int, last_time int) int {
	if index == len(time) {
		return 0
	}

	current_time += time[index]
	// choose index or not
	new1 := dfs(time, index+1, time[index], time[index]) + time[index]
	if new1 > current_time {
		return new1
	}

	new2 := dfs(time, index+1, current_time+last_time, last_time) + last_time

	max(new1, new2)

}

func Solution(battery []int, time []int, b int) int {

	dfs(time, 0, time[0], time[0])

	return result
}
