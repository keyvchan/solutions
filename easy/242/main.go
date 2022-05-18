package main

import "fmt"

func main() {
	fmt.Println(isAnagram("aacc", "ccac"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := map[rune]int{}

	for _, b := range s {
		set[b] += 1
	}
	fmt.Println(set)

	for _, b := range t {
		fmt.Println(set)
		if _, ok := set[b]; !ok {
			return false
		} else {
			set[b] -= 1
			if set[b] == 0 {
				delete(set, b)
			}
		}
	}

	return true
}
