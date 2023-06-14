package main

import (
	"fmt"
)

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
}

func buildArray(target []int, n int) []string {

	// there is a pointer indexed to target
	pointer := 0

	opereations := []string{}

	// start take number from stream
	for i := 1; i <= n; i++ {
		if pointer == len(target) {
			break
		}
		// check current number
		if target[pointer] == i {
			pointer++
			opereations = append(opereations, "Push")
		} else {
			opereations = append(opereations, "Push")
			opereations = append(opereations, "Pop")
		}

	}

	return opereations
}
