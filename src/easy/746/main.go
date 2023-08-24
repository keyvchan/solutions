package main

func main() {

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {

	min_costs_map := map[int]int{}

	for i := len(cost) - 1; i >= 0; i-- {
		if i == len(cost)-1 {
			min_costs_map[i] = cost[i]
		} else if i == len(cost)-2 {
			min_costs_map[i] = cost[i] + min_costs_map[i+1]
		} else {
			min_costs_map[i] = cost[i] + min(min_costs_map[i+1], min_costs_map[i+2])
		}
	}

	return min(min_costs_map[0], min_costs_map[1])

}
