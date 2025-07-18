package main

import "fmt"

func isAnagram(s string, t string) bool {
	// if lenght is different, then it is not an anagram
	if len(s) != len(t) {
		return false
	}

	// at least s is length
	charMapS := make(map[rune]int, len(s))

	for _, char := range s {
		charMapS[char]++
	}

	for _, char := range t {
		if charMapS[char] == 0 {
			return false
		}
		// char is in s
		charMapS[char]--
		if charMapS[char] == 0 {
			// char in s is used up
			delete(charMapS, char)
		}
	}

	return len(charMapS) == 0
}

func main() {
	for _, test := range []struct {
		s string
		t string
	}{
		{"anagram", "nagaram"},
		{"rat", "car"},
	} {
		fmt.Println(isAnagram(test.s, test.t))
	}
}
