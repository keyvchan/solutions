package main

import (
	"fmt"
)

func makeFancyString(s string) string {
	if s == "" {
		return ""
	}
	result := []byte{}
	lastChar := s[0]
	count := 0

	for i := 0; i < len(s); {
		if lastChar == s[i] {
			count++
			i++
			if count < 3 {
				result = append(result, lastChar)
			}
		} else {
			count = 0
			lastChar = s[i]
		}
	}

	return string(result)

}

func main() {
	fmt.Println(makeFancyString("abc"))
	fmt.Println(makeFancyString("aaabaaaa"))
	fmt.Println(makeFancyString("leeeetcode"))
}
