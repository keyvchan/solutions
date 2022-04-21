package main

import "fmt"

func main() {

	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
	fmt.Println(isSubsequence("acb", "ahbgdc"))

}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}

	for i, j := 0, 0; j < len(t); {

		if s[i] == t[j] {
			i++
			j++
		}
		if i >= len(s) {
			return true
		}
		if j >= len(t) {
			continue
		}

		if s[i] != t[j] {
			j++
		}

	}

	return false
}
