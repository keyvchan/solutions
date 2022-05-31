package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
}

func checkPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {

	last := string(s[len(s)-1])
	subset := []string{last}
	for i := len(s) - 2; i >= 0; i-- {

		for j := 0; j < len(subset); j++ {
			// update value of subset
			subset[j] = string(s[i]) + subset[j]
		}
		subset = append(subset, string(s[i]))
		// get all candidates
		candidates := []string{last}
		candidates = append(candidates, subset...)
		palindrome_len := 0
		for _, candidate := range candidates {
			// check if candidate is palindrome
			if checkPalindrome(candidate) {
				if len(candidate) > palindrome_len {
					palindrome_len = len(candidate)
					last = candidate
				}
			}
		}

	}

	return last
}
