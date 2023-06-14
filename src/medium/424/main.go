package main

import (
	"fmt"
)

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
	fmt.Println(characterReplacement("ABAA", 0))
}

func max(m map[byte]int) int {
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
func characterReplacement(s string, k int) int {

	char_map := make(map[byte]int, 26)
	if len(s) == 0 {
		return 0
	}
	char_map[s[0]] = 1
	max_window := 0

	// window_len - max <= k
	for i, j := 0, 0; ; {
		// get window_len
		window_len := j - i + 1
		// get max count
		max_freq := max(char_map)

		if window_len-max_freq <= k {
			if window_len > max_window {
				max_window = window_len
			}
			// extend window_len
			j += 1
			if j == len(s) {
				break
			}
			// update the char_map
			char_map[s[j]] += 1
		} else {
			char_map[s[i]] -= 1
			i += 1
			if i == len(s) {
				break
			}
			// update the char_map
		}

	}
	return max_window

}
