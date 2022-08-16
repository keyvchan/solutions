package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {

	var n, t int

	for {
		_, err := fmt.Scanln(&n, &t)
		if err == io.EOF {
			break
		}
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}
		result := solution(t, a)
		fmt.Println(result)

	}

}

func solution(t int, a []int) int {
	// for every order, check if it could be timeout
	sort.Ints(a)
	current_time := 0
	counts := 0
	for i := 0; i < len(a); i++ {
		// check if the first one if timeout
		if a[i]-current_time >= t {
			// we have enouth time to finish the order
			// current_time
			current_time += t
		} else {
			counts++
		}

	}
	return counts

}
