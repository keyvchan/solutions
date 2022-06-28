package main

import "fmt"

func main() {
	// fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	// fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))
}

type Prefix struct {
	last_index int
	word       string
}

func check(s string, index int, next_word string, dict map[string][]string, caching map[Prefix]bool) bool {
	// fmt.Println(s, index, next_word)
	fmt.Println(index, next_word)
	// base case
	if _, ok := caching[Prefix{index, next_word}]; ok {
		return caching[Prefix{index, next_word}]
	}
	if index+len(next_word) >= len(s) {
		return s[index:] == next_word
	}
	if s[index:index+len(next_word)] != next_word {
		return false
	}

	result := false
	for _, word := range dict[string(s[index+len(next_word)])] {
		temp := check(s, index+len(next_word), word, dict, caching)
		result = result || temp
	}
	caching[Prefix{index, next_word}] = result

	return result
}

func wordBreak(s string, wordDict []string) bool {
	// create a map to store first letter map to entire words
	dict := map[string][]string{}
	for _, word := range wordDict {
		dict[string(word[0])] = append(dict[string(word[0])], word)
	}
	// add worddict to map for quick word check
	wordDictMap := map[string]bool{}
	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	caching := map[Prefix]bool{}
	result := false
	for _, word := range dict[string(s[0])] {
		result = result || check(s, 0, word, dict, caching)
	}

	return result
}
