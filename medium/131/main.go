package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("abbab"))
}

func backtracking(s string, rest string, current_string []string, result *[][]string, all_palin map[string]bool) {
	if all_palin[s] {
		current_string = append(current_string, s)
	} else {
		return
	}
	if len(rest) == 0 {
		*result = append(*result, current_string)
		return
	}

	// backtracking all substrings of s
	for i := range rest {
		// copy current_string to avoid changing the original
		new_current_string := make([]string, len(current_string))
		copy(new_current_string, current_string)
		backtracking(rest[:i+1], rest[i+1:], new_current_string, result, all_palin)
	}
}

func checkPalindrome(s string, i int, j int, all_palin map[string]bool) {
	for i >= 0 && j < len(s) {
		fmt.Println(i, j)
		if s[i] == s[j] {
			all_palin[s[i:j+1]] = true
		} else {
			break
		}
		i--
		j++
	}
}

func partition(s string) [][]string {

	all_palin := map[string]bool{}

	// find all palindromes
	for i := 0; i < len(s); i++ {
		// odd
		checkPalindrome(s, i, i, all_palin)
		checkPalindrome(s, i, i+1, all_palin)

	}
	fmt.Println(all_palin)

	result := [][]string{}
	// check every substring is a palindrome
	// backtracking all substrings of s
	current_string := []string{}
	for i := range s {
		// copy current_string to avoid changing the original
		new_current_string := make([]string, len(current_string))
		copy(new_current_string, current_string)
		backtracking(s[:i+1], s[i+1:], new_current_string, &result, all_palin)
	}

	return result
}
