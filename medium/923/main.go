package main

import "sort"

func main() {

}

func threeSumMulti(arr []int, target int) int {

	sort.Ints(arr)

	counts := make(map[int]int)

	// find every number occur times
	for _, v := range arr {
		_, ok := counts[v]
		if !ok {
			counts[v] = 1
		} else {
			counts[v] = counts[v] + 1
		}
		if v > target {
			break
		}
	}

	return 1
}
