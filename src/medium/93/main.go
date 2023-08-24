package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("101023"))
}

func push_to_temp(s string, index int, count int) []string {
	if count == 3 {
		if is_valid(s[index:]) {
			return []string{s[index:]}
		} else {
			return []string{}
		}
	}

	result := []string{}
	max_index := 0
	if index+3 > len(s)-1 {
		max_index = len(s) - 1
	} else {
		max_index = index + 3
	}

	for i := index; i < max_index; i++ {
		new_string := s[index : i+1]
		// check new_string is valid
		if !is_valid(new_string) {
			continue
		}
		rests := push_to_temp(s, i+1, count+1)
		for _, rest := range rests {
			final_string := new_string + "." + rest
			result = append(result, final_string)
		}
	}
	return result
}

func is_valid(s string) bool {
	if len(s) <= 0 || len(s) > 3 {
		return false
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i < 0 || i > 255 {
		return false
	}
	if s[0] == '0' && len(s) > 1 {
		return false
	}
	return true
}

func restoreIpAddresses(s string) []string {

	return push_to_temp(s, 0, 0)

}
