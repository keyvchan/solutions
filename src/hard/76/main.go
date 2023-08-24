package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("ab", "a"))
	fmt.Println(minWindow("ab", "b"))
	fmt.Println(minWindow("bdab", "ab"))
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
	fmt.Println(minWindow("acbbaca", "baca"))
}

func minWindow(s string, t string) string {

	t_map := map[byte]int{}

	for _, v := range t {
		t_map[byte(v)] += 1
	}

	s_map := map[byte]int{}

	win_len := 0
	win_start := 0
	win_end := 0
	for i, j := 0, 0; ; {
		fmt.Println(i, j)
		if reflect.DeepEqual(s_map, t_map) {
			fmt.Println(s_map, t_map)
			// update window length
			if j-i+1 < win_len || win_len == 0 {
				win_len = j - i + 1
				win_start = i
				win_end = j
			}
			// minus counts of the first char in the window
			s_map[s[i]] -= 1
			i++
			// shirnk window from left
			for i < len(s) {
				if _, ok := t_map[s[i]]; ok {
					if reflect.DeepEqual(s_map, t_map) {
						if j-i+1 < win_len || win_len == 0 {
							win_len = j - i + 1
							win_start = i
							win_end = j - 1
						}
					}
					break
				} else {
					i++
				}
			}
		} else {
			// shirnk window from left
			if len(s_map) == 0 {
				// i follows j
				i = j
			}
			// expand window from right
			if _, ok := t_map[s[j]]; ok {
				s_map[s[j]] += 1
			}
			fmt.Println(s_map, t_map)
			if s_map[s[j]] > t_map[s[j]] {
				fmt.Println(">")
				// if s[j] != s[i], ignore
				if s[j] == s[i] {
					// set i to j, and reset s_map
					i = j
					s_map = map[byte]int{}
					s_map[s[i]] += 1
				} else {
					s_map[s[j]] -= 1
				}
			}
			j++
		}
		if j == len(s) || i == len(s) {
			// final check
			if reflect.DeepEqual(s_map, t_map) {
				if j-i+1 < win_len || win_len == 0 {
					win_len = j - i + 1
					win_start = i
					win_end = j
				}
			}
			break
		}
	}

	if win_len == 0 {
		return ""
	} else {
		return s[win_start:win_end]
	}

}
