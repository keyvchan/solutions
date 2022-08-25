package main

import "fmt"

func main() {
	fmt.Println(pushDominoes(".L.R...LR..L.."))

}

func pushDominoes(dominoes string) string {
	// right := make([]int, len(dominoes))
	left := make([]int, len(dominoes))
	right := make([]int, len(dominoes))

	for i := 1; i < len(dominoes); i++ {
		// check i-1
		if dominoes[i] == '.' {
			if dominoes[i-1] == 'R' || (dominoes[i-1] == '.' && left[i-1] != 0) {
				left[i] = left[i-1] + 1
			}
		}
	}
	for i := len(dominoes) - 2; i >= 0; i-- {
		// check i-1
		if dominoes[i] == '.' {
			if dominoes[i+1] == 'L' || (dominoes[i+1] == '.' && right[i+1] != 0) {
				right[i] = right[i+1] + 1
			}
		}
	}

	res := make([]byte, len(dominoes))

	// compare each
	for i := 0; i < len(dominoes); i++ {
		// left and right both 0
		if left[i] == 0 && right[i] == 0 {
			res[i] = dominoes[i]
			continue
		}
		if left[i] == 0 && right[i] != 0 {
			res[i] = 'L'
			continue
		}
		if right[i] == 0 && left[i] != 0 {
			res[i] = 'R'
			continue
		}
		if left[i] == right[i] {
			res[i] = '.'
			continue
		}
		if left[i] < right[i] {
			res[i] = 'R'
		} else {
			res[i] = 'L'
		}
	}

	return string(res)
}
