package main

import (
	"fmt"
)

func main() {
	fmt.Println(partition("aab"))
}

func isPalin(s string) bool {
	reverse := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	new_string := reverse(s)
	if new_string == s {
		return true
	} else {
		return false
	}
}

func backtracking(s string, current_string string, all_palin *[][]string) {

	if isPalin(current_string) {
		*all_palin = append(*all_palin, []string{current_string})
	}

	for i, v := range s {
		current_string += string(v)
		backtracking(s[i+1:], current_string, all_palin)
		current_string = current_string[:len(current_string)-1]
	}

}

func partition(s string) [][]string {

	all_palin := [][]string{}

	backtracking(s, "", &all_palin)

	return all_palin
}
