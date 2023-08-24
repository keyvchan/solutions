package main

import "fmt"

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
}

func partitionLabels(s string) []int {

	// count words frequency
	chars := map[rune]int{}

	for _, c := range s {
		chars[c]++
	}

	encounters := map[rune]bool{}

	last := 0
	result := []int{}
	for i, char := range s {

		if !encounters[char] {
			encounters[char] = true
		}
		chars[char]--
		if chars[char] == 0 {
			delete(encounters, char)
		}
		if len(encounters) == 0 {
			// this is a part
			result = append(result, i-last+1)
			last = i + 1
		}

	}

	return result
}
