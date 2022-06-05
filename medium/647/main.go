package main

func main() {

}

func countSubstrings(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			count++
			left--
			right++
		}
		left, right = i, i+1
		for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
			count++
			left--
			right++
		}
	}
	return count

}
