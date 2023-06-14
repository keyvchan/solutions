package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"", ""}))
	fmt.Println(groupAnagrams([]string{"h", "h", "h", "h"}))
	fmt.Println(groupAnagrams([]string{"ac", "c"}))
}

func groupAnagrams(strs []string) [][]string {
	// generate all anagram map
	result := [][]string{}
	result_map := map[[26]int][]string{}
	for _, str := range strs {
		new_array := [26]int{}
		for _, r := range str {
			new_array[r-'a']++
		}
		result_map[new_array] = append(result_map[new_array], str)
	}

	for _, v := range result_map {
		result = append(result, v)
	}
	return result
}
