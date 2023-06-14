package main

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}

func isSubsequence(s string, t string) bool {
	if len(t) > len(s) {
		return false
	}
	if s == t {
		return true
	}
	// use two pointer
	i, j := 0, 0
	for {
		if i == len(s) || j == len(t) {
			break
		}
		// i tracking the index of s, j tracking the index of t
		if s[i] == t[j] {
			i++
			j++
		} else {
			i++
		}

	}
	return j == len(t)

}

func numMatchingSubseq(s string, words []string) int {

	total := 0
	for _, word := range words {
		if isSubsequence(s, word) {
			total++
		}
	}
	return total

}
