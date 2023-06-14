package main

func main() {

}

func longestPalindrome(s string) int {

	// count words frequency
	freq := map[byte]int{}

	for _, c := range s {
		freq[byte(c)]++
	}

	// the final result is made up of the following parts
	// 1. even number of words
	// 2. the odd number of words with the largest frequency
	// 3. all other odd number of words - 1

	// 1. even number of words
	even := 0
	odd := 0
	odd_pairs := 0

	largestOdd := 0

	for _, v := range freq {
		if v%2 == 0 {
			even += v
		} else {
			odd += v
			odd_pairs++
			if v > largestOdd {
				largestOdd = v
			}
		}
	}
	if odd_pairs == 0 {
		return even + largestOdd
	}

	return (even) + (largestOdd) + (odd - largestOdd - odd_pairs + 1)

}
