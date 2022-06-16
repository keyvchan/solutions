package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
}

func longestPalindrome(s string) string {
	get_palindrome := func(s string, left, right int, length *int, sub *string) {
		for {
			if left < 0 || right >= len(s) || s[left] != s[right] {
				if *length < right-left {
					*length = right - left
					*sub = s[left+1 : right]
				}
				break
			}
			if s[left] == s[right] {
				// update length
				if right-left+1 > *length {
					*length = right - left + 1
				}
			}
			left--
			right++
		}

	}
	length := 0
	sub := ""
	for i := 0; i < len(s); i++ {
		left, right := i, i
		get_palindrome(s, left, right, &length, &sub)

		left, right = i, i+1
		get_palindrome(s, left, right, &length, &sub)

	}
	return sub
}
