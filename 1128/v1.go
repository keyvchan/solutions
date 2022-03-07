package main

import "fmt"

func main() {
	dominoes1 := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	dominoes2 := [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}

	fmt.Println(numEquivDominoPairs(dominoes1))
	fmt.Println(numEquivDominoPairs(dominoes2))

}

func numEquivDominoPairs(dominoes [][]int) int {
	sum := 0

	for i := 0; i < len(dominoes); i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) || (dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0]) {
				sum += 1
			}
		}
	}

	return sum
}
