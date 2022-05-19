package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {

	// remove all non-alphanumeric characters
	s = strings.ToLower(s)
	new_string := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			// append to new string
			new_string += string(s[i])
		}

	}

	for i, j := 0, len(new_string)-1; i < j; i, j = i+1, j-1 {
		if new_string[i] != new_string[j] {
			return false
		}
	}

	return true
}
