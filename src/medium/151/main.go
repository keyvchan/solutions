package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	S := strings.Split(s, " ")
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		S[i], S[j] = S[j], S[i]
	}

	result := ""
	for _, ss := range S {
		ss = strings.TrimSpace(ss)
		if len(ss) == 0 {
			continue
		}
		result += ss
		result += " "
	}

	result = strings.TrimSpace(result)

	return result
}
