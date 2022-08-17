package main

import (
	"fmt"
)

func main() {
	fmt.Println(minDeletions("aabb"))
	fmt.Println(minDeletions("abb"))
	fmt.Println(minDeletions("aaabbbcc"))
	fmt.Println(minDeletions("ceabaacb"))
	fmt.Println(minDeletions("ee"))
}

func minDeletions(s string) int {
	// counts char frequency

	char_freq := map[rune]int{}

	for _, char := range s {
		char_freq[char]++
	}

	// the max frequency is the string length
	freqs := make([]int, len(s)+1)
	for _, freq := range char_freq {
		freqs[freq]++
	}

	counts := 0
	// reverse loop the whole array
	for i := len(freqs) - 1; i > 0; i-- {
		if freqs[i] > 1 {
			counts += freqs[i] - 1
			freqs[i-1] += freqs[i] - 1
		}
	}
	return counts

}
