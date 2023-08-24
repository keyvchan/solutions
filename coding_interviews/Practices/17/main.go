package main

import "math"

func main() {

}

func printNumbers(n int) []int {
	// calcuate the max number with n width
	max := math.Pow10(n)

	reuslt := make([]int, 0, int(max)-1)
	for i := 1; i < int(max); i++ {
    reuslt = append(reuslt, i)
	}

	return reuslt
}
