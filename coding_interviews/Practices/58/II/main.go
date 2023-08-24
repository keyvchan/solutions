package main

func main() {

}

func reverseLeftWords(s string, n int) string {
	left := s[:n]
	right := s[n:]

	result := right + left
	return result

}
