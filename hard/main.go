package main

import "fmt"

func main() {

	fmt.Println(kSimilarity("ab", "ba"))
	fmt.Println(kSimilarity("abc", "bca"))
	fmt.Println(kSimilarity("abcd", "badc"))
	fmt.Println(kSimilarity("aabbcc", "aabbcc"))
	fmt.Println(kSimilarity("abac", "baca"))
	fmt.Println(kSimilarity("aabc", "abca"))
	fmt.Println(kSimilarity("abccaacceecdeea", "bcaacceeccdeaae"))

}

func kSimilarity(s1 string, s2 string) int {

	return 0

	// differents := 0
	// for i := 0; i < len(s1); i++ {
	// 	if s1[i] != s2[i] {
	// 		differents = differents + 1
	// 	}
	// }

	// if differents%2 == 0 {
	// 	return differents / 2
	// } else {
	// 	return differents/2 + 1
	// }
}
