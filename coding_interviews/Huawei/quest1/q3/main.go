package main

import (
	"fmt"
	"io"
)

func main() {

	var D int
	var N int
	for {
		_, err := fmt.Scan(&D)
		if err == io.EOF {
			break
		}
		fmt.Scan(&N)
		station := make([]Station, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&station[i].distance)
			fmt.Scan(&station[i].line)
		}
		fmt.Println(Solution(D, station))
	}

}

type Station struct {
	distance int
	line     int
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func dfs(distance int, current_station int, current_battery int, station []Station, memo map[Key]int) int {
	if current_station == len(station)-1 {
		// we cahek if we have enough battery to reach the end
		if current_battery >= distance-station[current_station].distance {
			return 0
		} else {
			// we need to charge
			// check full battery can reach the end
			if distance-station[current_station].distance > 1000 {
				// we can't reach the end
				return 100000
			} else {
				// we need to charge
				return station[current_station].line + 1
			}
		}
	}

	if val, ok := memo[Key{current_station, current_battery}]; ok {
		return val
	}

	min_time := 100000

	to_next := station[current_station+1].distance - station[current_station].distance

	if to_next > 1000 {
		return 100000
	}

	min_time = min(min_time, dfs(distance, current_station+1, current_battery-to_next, station, memo))
	min_time = min(min_time, dfs(distance, current_station+1, 1000-to_next, station, memo)+station[current_station].line+1)

	memo[Key{current_station, current_battery}] = min_time

	return min_time

}

type Key struct {
	current_station int
	current_battery int
}

func Solution(distance int, station []Station) int {
	if len(station) == 0 {
		if distance > 1000 {
			return -1
		}
	}
	// for every station, we can stop then charge it
	// check if we can reach 0 station
	if station[0].distance > 1000 {
		return -1
	}

	memo := map[Key]int{}

	result := dfs(distance, 0, 1000-station[0].distance, station, memo) + distance/100
	if result > 100000 {
		return -1
	}
	return result
}
