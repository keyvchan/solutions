package main

import "strconv"

func main() {

}

func secondHighest(s string) int {

	// first highest
	var first = -1

	// second highest
	var second = -1

	// loop through string
	for _, s := range s {
		num, err := strconv.Atoi(string(s))
		if err != nil {
			// it's not a number
			continue
		}
		// its a number, check if it's higher than first
		if num > first {
			// it is, set first to num
			second = first
			first = num
		} else {
			if num == first {
				continue
			}
			// it's not, check if it's higher than second
			if num > second {
				// it is, set second to num
				second = num
			}
			// it's not, continue
		}

	}
	if second == first {
		return -1
	}

	return second
}
