package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("ckilbkd"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
}

func lengthOfLongestSubstring(s string) int {
	distinct_map := make(map[byte]int)
	max := 0
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	for i, j := 0, 0; ; {
		if _, ok := distinct_map[s[j]]; ok {
			// find the char index, delete all before it
			if max < len(distinct_map) {
				max = len(distinct_map)
			}
			index := distinct_map[s[j]]
			for k, v := range distinct_map {
				if v <= index {
					delete(distinct_map, k)
				}
			}
			distinct_map[s[j]] = j
			i = index
			j++

		} else {
			distinct_map[s[j]] = j
			j++
		}
		if i == len(s) || j == len(s) {
			break
		}
	}

	if max < len(distinct_map) {
		max = len(distinct_map)
	}

	return max

}
