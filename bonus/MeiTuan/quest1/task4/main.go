package main

import (
	"fmt"
	"io"
)

func main() {
	var n int

	for {
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		fmt.Print(solution(a))
	}

}

func solution(a []int) int {
	// a[i] + a[k] = 3a[j]
	// use a map to store all number
	if len(a) < 3 {
		return 0
	}
	nums := map[int]bool{}

	for i := 0; i < len(a); i++ {
		nums[a[i]] = true
	}

	counts := 0
	// check every tuple
	for i := 0; i < len(a)-2; i++ {
		for j := i + 1; j < len(a)-1; j++ {
			if nums[3*a[j]-a[i]] {
				counts++
			}
		}
	}
	return counts
}
