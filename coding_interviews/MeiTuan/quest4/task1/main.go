package main

import (
	"fmt"
	"io"
)

func main() {
	m, n := 0, 0
	var S string
	var s string

	for {
		_, err := fmt.Scan(&m, &n)
		if err == io.EOF {
			break
		}

		fmt.Scan(&S)
		fmt.Scan(&s)
		fmt.Println(Solution(S, s))
	}

}

func Solution(S string, s string) int {
	// from start of S, check if s is a substring

	j, k := 0, 0
	counts := 0

	for i := 0; i < len(S); i++ {
		j = i
		for {
			if j >= len(S) || k >= len(s) {
				break
			}
			if j+k >= len(S) {
				break
			}
			// remain of S is shorter than s
			if s[k] == '*' {
				// * can match any character
				j++
				k++
				continue
			}
			if S[j+k] == s[k] {
				j++
				k++
			} else {
				// mismatch
				// we check if j == len(s)
				break
			}
		}
		if k == len(s)-1 {
			counts++
		}
		k = 0
	}

	return counts

}
