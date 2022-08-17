package main

import "fmt"

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))

}

func uniqueMorseRepresentations(words []string) int {
	morse := [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	// use a hashmap to store all unique morse representations
	m := map[string]bool{}
	for _, word := range words {
		morseWord := ""
		for _, char := range word {
			morseWord += morse[char-'a']
		}
		m[morseWord] = true
	}

	return len(m)
}
