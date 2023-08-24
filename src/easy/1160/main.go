package main

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
	fmt.Println(countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr"))
}

func countCharacters(words []string, chars string) int {
	// chars map

	chars_map := map[string]int{}

	for _, char := range chars {
		chars_map[string(char)]++
	}
	sum := 0

	// check every word
	for _, word := range words {
		// word map
		word_map := map[string]int{}
		for _, char := range word {
			word_map[string(char)]++
		}
		sum += len(word)

		for key, value := range word_map {
			if chars_map[key] < value {
				sum -= len(word)
				break
			}
		}
	}
	return sum

}
