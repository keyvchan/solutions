package main

func main() {

}

func findTheDifference(s string, t string) byte {
	// count word frequency
	s_words := map[rune]int{}
	for _, ss := range s {
		s_words[ss]++
	}

	for _, tt := range t {
		s_words[tt]--
		if s_words[tt] < 0 {
			return byte(tt)
		}
	}

	return byte(' ')
}
