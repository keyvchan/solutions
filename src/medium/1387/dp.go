package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(getKth(1, 22, 2))
	fmt.Println(getKth(7, 11, 4))
	fmt.Println(getKth(1, 1000, 777))
}

func getKth(lo int, hi int, k int) int {
	dp := map[int]int{}
	// from lo to hi, store the power
	for i := 1; i <= hi; i++ {
		temp := i
		counts := 0
		for temp != 1 {
			if _, ok := dp[temp]; ok {
				fmt.Println("hit", temp)
				counts = counts + dp[temp]
				break
			}
			if temp%2 == 0 {
				temp = temp / 2
			} else {
				temp = temp*3 + 1
			}
			counts++
		}
		dp[i] = counts
	}

	// create and sort
	var arr []int
	for i := lo; i <= hi; i++ {
		arr = append(arr, i)
	}
	sort.Slice(arr, func(i, j int) bool {
		if dp[arr[i]] == dp[arr[j]] {
			return arr[i] < arr[j]
		}
		return dp[arr[i]] < dp[arr[j]]
	})
	return arr[k-1]
}
