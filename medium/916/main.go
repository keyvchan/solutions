package main

import "fmt"

func main() {
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
	fmt.Println(wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
}

func isUniversal(word1 string, words2 map[rune]int) bool {
	// check letter frequency of word1
	word1Freq := make(map[rune]int)
	for _, char := range word1 {
		word1Freq[char]++
	}
	for char, count := range words2 {
		word1Freq[char] -= count
		if word1Freq[char] < 0 {
			return false
		}
	}

	return true
}

func wordSubsets(words1 []string, words2 []string) []string {
	result := []string{}

	// construct a map of word2s letter max frequencies
	word2Freq := make(map[rune]int)
	for _, word2 := range words2 {
		temp_map := make(map[rune]int)
		for _, char := range word2 {
			temp_map[char]++
		}
		for char, count := range temp_map {
			if count > word2Freq[char] {
				word2Freq[char] = count
			}
		}
	}
	// for every word in words1, check if its universal
	for _, word := range words1 {
		if isUniversal(word, word2Freq) {
			result = append(result, word)
		}
	}
	return result
}
