package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}

func isAlphanumeric(s byte) bool {
	if s >= 'a' && s <= 'z' || s >= '0' && s <= '9' {
		return true
	}
	return false
}

func isPalindrome(s string) bool {

	// remove all non-alphanumeric characters
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; {
		// check s[i] is alphanumeric
		if !isAlphanumeric(s[i]) {
			i += 1
			continue
		}
		if !isAlphanumeric(s[j]) {
			j -= 1
			continue
		}

		if s[i] != s[j] {
			return false
		} else {
			i += 1
			j -= 1
		}
	}

	return true
}
