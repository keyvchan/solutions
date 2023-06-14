package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hasAllCodes("00110", 2))
}

func hasAllCodes(s string, k int) bool {
	// counts the number of s distinct substrings of length k

	if k > len(s) {
		return false
	}

	left := 0
	right := k

	counts_map := map[string]bool{}
	for {
		current := s[left:right]
		if !counts_map[current] {
			// current is a distinct substring
			counts_map[current] = true
		}
		left += 1
		right += 1
		if right == len(s)+1 {
			break
		}
	}
	counts := len(counts_map)
	fmt.Println(counts_map)

	if float64(counts) == math.Pow(2, float64(k)) {
		return true
	} else {
		return false
	}
}
