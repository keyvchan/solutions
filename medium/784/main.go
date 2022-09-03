package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
}

func backtracking(s string, index int, current string, result *[]string) {
	if index == len(s) {
		*result = append(*result, current)
		return
	}
	if s[index] >= '0' && s[index] <= '9' {
		// we only have one choice
		backtracking(s, index+1, current+string(s[index]), result)
	} else {
		// we have two choices
		lower := strings.ToLower(string(s[index]))
		upper := strings.ToUpper(string(s[index]))
		backtracking(s, index+1, current+lower, result)
		backtracking(s, index+1, current+upper, result)
	}

}

func letterCasePermutation(s string) []string {
	result := []string{}

	backtracking(s, 0, "", &result)

	return result
}
