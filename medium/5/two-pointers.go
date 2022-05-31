package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
}

func longestPalindrome(s string) string {
	length := 0
	sub := ""
	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// update length
			if right-left+1 > length {
				length = right - left + 1
				sub = s[left : right+1]
			}
			left--
			right++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			// update length
			if right-left+1 > length {
				length = right - left + 1
				sub = s[left : right+1]
			}
			left--
			right++
		}

	}
	return sub
}
