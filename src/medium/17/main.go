package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

var digits_map = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func combination(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return digits_map[string(digits[0])]
	}
	result := []string{}

	runes := digits_map[string(digits[0])]
	sub_strings := combination(digits[1:])
	for _, r := range runes {
		for _, sub_string := range sub_strings {
			new_string := string(r) + sub_string
			result = append(result, new_string)
		}
	}

	return result
}

func letterCombinations(digits string) []string {

	return combination(digits)

}
