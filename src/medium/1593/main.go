package main

import "fmt"

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
	fmt.Println(maxUniqueSplit("aba"))
	fmt.Println(maxUniqueSplit("aa"))
	fmt.Println(maxUniqueSplit("wwwzfvedwfvhsww"))
}
func backtracking(s string, index int, current string, result map[string]bool, counts *int) {
	if index == len(s) {
		if len(result) > *counts {
			*counts = len(result)
		}
		return
	}
	// we could split here or not
	current += string(s[index])
	if _, ok := result[current]; !ok {
		// split here
		result[current] = true
		backtracking(s, index+1, "", result, counts)
		delete(result, current)

		// or not
		backtracking(s, index+1, current, result, counts)
	} else {
		// not split here
		backtracking(s, index+1, current, result, counts)
	}
}

func maxUniqueSplit(s string) int {
	// use a hash set to store the unique substrings
	counts := 0
	result := map[string]bool{}
	backtracking(s, 0, "", result, &counts)

	return counts
}
