package main

func main() {

}

func firstUniqChar(s string) byte {
	// count word frequency
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// find the first unique character
	for i := 0; i < len(s); i++ {
		if freq[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}
