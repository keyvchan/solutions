package main

import (
	"fmt"
)

func main() {
	fmt.Println(brokenCalc(2, 3))
	fmt.Println(brokenCalc(5, 8))
	fmt.Println(brokenCalc(3, 10))
}

func brokenCalc(startValue int, target int) int {
	steps := 0
	// from target to startValue

	for target > startValue {
		if target%2 == 0 {
			target = target / 2
			steps++
		} else {
			target = target + 1
			steps++
		}
	}

	return steps + startValue - target
}
