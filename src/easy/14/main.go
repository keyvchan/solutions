package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		if i == len(strs[0]) {
			return strs[0][:i]
		}
		current_char := strs[0][i]
		// check all the str if str[i] is the same
		// check if index out of range
		for _, str := range strs {
			if i >= len(str) {
				return str[:i]
			}
			if str[i] != current_char {
				// if not the same, return the prefix
				return str[:i]
			}
		}

	}

}
