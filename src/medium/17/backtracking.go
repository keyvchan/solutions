package main

func main() {

}

var digits_map = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func backtracking(digits string, current_digits string, result *[]string) {
	if len(digits) == 0 {
		*result = append(*result, current_digits)
		return
	}
	// take digits[0]
	for _, letter := range digits_map[digits[0]] {
		// every letter in the map
		backtracking(digits[1:], current_digits+letter, result)
	}

}

func letterCombinations(digits string) []string {

	result := []string{}
	backtracking(digits, "", &result)
	return result
}
