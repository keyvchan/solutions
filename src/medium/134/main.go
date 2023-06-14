package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}
func canCompleteCircuit(gas []int, cost []int) int {
	// -2, -2, -2, 3, 3
	sum_gas := 0
	sum_cost := 0
	for i := 0; i < len(gas); i++ {
		sum_gas += gas[i]
		sum_cost += cost[i]
	}
	if sum_gas < sum_cost {
		return -1
	}

	total := 0
	current := 0
	// start from 0, check every possible start point
	for i := 0; i < len(gas); i++ {
		fmt.Println(total, current, gas[i], cost[i])
		if total == 0 {
			current = i
		}
		total += gas[i] - cost[i]

		if total < 0 {
			total = 0
		}
	}
	if total < 0 {
		// end of the loop, total is still negative
		return -1
	}

	return current
}
