package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
	fmt.Println(validPalindrome("bddb"))
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))
	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}
func isPalindrome(s string, flag bool) bool {
	i, j := 0, len(s)-1
	deleted := flag
	for {
		if i >= j {
			break
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			if deleted {
				return false
			} else {
				deleted = true
				return isPalindrome(s[i+1:j+1], true) || isPalindrome(s[i:j], true)
			}
		}
	}
	return true
}

func validPalindrome(s string) bool {
	return isPalindrome(s, false)
}
