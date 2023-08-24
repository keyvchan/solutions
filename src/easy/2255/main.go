package main

import "fmt"

func main() {
	fmt.Println(countPrefixes([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))
}

func countPrefixes(words []string, s string) int {
	count := 0
	for _, word := range words {
		if len(word) > len(s) {
			continue
		}
		// check prefix
		if s[:len(word)] == word {
			count++
		}
	}
	return count

}
