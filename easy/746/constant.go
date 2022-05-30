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

	last := cost[len(cost)-1]
	secondLast := cost[len(cost)-2]

	for i := len(cost) - 3; i >= 0; i -= 2 {
		last = cost[i] + min(last, secondLast)
		secondLast = cost[i-1] + min(last, secondLast)
	}

	return min(last, secondLast)
}
