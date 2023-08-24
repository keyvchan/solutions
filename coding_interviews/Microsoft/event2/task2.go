package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution([]int{1, 1, 2}, []int{3, 2, 3}))
	fmt.Println(Solution([]int{1, 1, 1}, []int{2, 2, 2}))
	fmt.Println(Solution([]int{1, 2, 3, 1, 2, 12, 8, 4}, []int{5, 10, 15, 2, 4, 15, 10, 5}))
}

type Fraction struct {
	numerator   int
	denominator int
}

func Solution(X []int, Y []int) int {
	// for every fraction, check if it can sum other fraction up to 1
	memo := map[Fraction]int{}
	counts := 0
	for i := 0; i < len(X); i++ {
		if _, ok := memo[Fraction{X[i], Y[i]}]; ok {
			counts += memo[Fraction{X[i], Y[i]}]
			continue
		}
		for j := 0; j < len(Y); j++ {
			if j == i {
				continue
			}
			// if denominator is not equal
			if Y[i] != Y[j] {
				if X[i]*Y[j]+X[j]*Y[i] == Y[i]*Y[j] {
					counts++
				}
			} else {
				if X[i]+X[j] == Y[i] {
					counts++
				}
			}
		}
		memo[Fraction{X[i], Y[i]}] = counts
	}
	return (counts / 2) % 1000000007

}
