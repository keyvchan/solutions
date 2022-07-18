package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkAlmostEquivalent("hello", "hella"))
	fmt.Println(checkAlmostEquivalent("hello", "hell"))
	fmt.Println(checkAlmostEquivalent("aaaa", "bccb"))
	fmt.Println(checkAlmostEquivalent("abcdeef", "abaaacc"))
	fmt.Println(checkAlmostEquivalent("cccddabba", "babababab"))
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	// count word frequency of word1
	var word1Freq = make(map[rune]int)
	var word2Freq = make(map[rune]int)
	for _, char := range word1 {
		word1Freq[char]++
	}

	for _, char := range word2 {
		word2Freq[char]++
	}
	// check from a-z
	for i := 'a'; i <= 'z'; i++ {
		// check difference between word1 and word2
		if math.Abs(float64(word1Freq[i]-word2Freq[i])) > 3 {
			return false
		}
	}
	return true

}
