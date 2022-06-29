package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
	fmt.Println(longestStrChain([]string{"qyssedya", "pabouk", "mjwdrbqwp", "vylodpmwp", "nfyqeowa", "pu", "paboukc", "qssedya", "lopmw", "nfyqowa", "vlodpmw", "mwdrqwp", "opmw", "qsda", "neo", "qyssedhyac", "pmw", "lodpmw", "mjwdrqwp", "eo", "nfqwa", "pabuk", "nfyqwa", "qssdya", "qsdya", "qyssedhya", "pabu", "nqwa", "pabqoukc", "pbu", "mw", "vlodpmwp", "x", "xr"}))
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isValid(nextword, word string) bool {
	// word delete a char and becomes nextword
	// the order cant be changed
	// try to delete every char in word
	for i := 0; i < len(word); i++ {
		if nextword == word[:i]+word[i+1:] {
			return true
		}
	}
	return false
}

func dfs(groups map[int][]string, word string, caching map[string]int) int {
	if _, ok := groups[len(word)-1]; !ok {
		return 0
	}
	if v, ok := caching[word]; ok {
		return v
	}

	max := 0
	for _, w := range groups[len(word)-1] {
		// check if the word is a valid next word
		if isValid(w, word) {
			max = maxInt(max, dfs(groups, w, caching)+1)
		}
	}
	caching[word] = max
	return max

}

func longestStrChain(words []string) int {

	// group words by length
	groups := map[int][]string{}

	for _, word := range words {
		groups[len(word)] = append(groups[len(word)], word)
	}

	// sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	caching := map[string]int{}

	max := 0
	for _, word := range words {
		// find the longest chain for each word
		max = maxInt(max, dfs(groups, word, caching)+1)
	}

	return max
}
