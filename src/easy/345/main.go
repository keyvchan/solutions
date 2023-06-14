package main

import "fmt"

func main() {

	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels(" "))
}

func reverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	ss := []byte(s)
	i := 0
	j := len(s) - 1
	for {
		if i > j {
			break
		}
		if !vowels[ss[i]] {
			i++
			continue
		}
		if !vowels[ss[j]] {
			j--
			continue
		}
		if vowels[ss[i]] && vowels[ss[j]] {
			ss[i], ss[j] = ss[j], ss[i]
			i++
			j--
		}
	}

	return string(ss)
}
