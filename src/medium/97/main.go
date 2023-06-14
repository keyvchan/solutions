package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("a", "", "a"))
	fmt.Println(isInterleave("a", "b", "a"))
	fmt.Println(isInterleave("aa", "ab", "abaa"))
}

func dfs(s1 string, s2 string, s3 string, s1_counts int, s2_counts int, cache map[Key]Value) Value {
	// fmt.Println(s1, s2, s3)
	if len(s3) == 0 {
		// check s1_counts and s2_counts
		if math.Abs(float64(s1_counts-s2_counts)) > 1 {
			return Value{false, s1_counts, s2_counts}
		} else {
			return Value{true, s1_counts, s2_counts}
		}
	}
	if len(s1) == 0 {
		// check s1_counts and s2_counts
		if s1_counts < s2_counts {
			return Value{false, s1_counts, s2_counts + 1}
		} else {
			if s1_counts-s2_counts != 1 && s1_counts-s2_counts != 0 {
				return Value{false, s1_counts, s2_counts}
			}
			if s2 == s3 {
				return Value{true, s1_counts, s2_counts + 1}
			} else {
				return Value{false, s1_counts, s2_counts}
			}
		}
	}
	if len(s2) == 0 {
		// fmt.Println(s1_counts, s2_counts)
		// fmt.Println("len(s2) == 0")
		// check s1_counts and s2_counts
		if s2_counts < s1_counts {
			// fmt.Println("s2_counts < s1_counts")
			return Value{false, s1_counts + 1, s2_counts}
		} else {
			if s2_counts-s1_counts != 1 && s2_counts-s1_counts != 0 {
				// fmt.Println("s2_counts-s1_counts != 1 && s2_counts-s1_counts != 0")
				return Value{false, s1_counts, s2_counts}
			}
			if s1 == s3 {
				return Value{true, s1_counts + 1, s2_counts}
			} else {
				return Value{false, s1_counts, s2_counts}
			}
		}
	}
	if _, ok := cache[Key{s1, s2, s3}]; ok {
		fmt.Println("cache hit")
		return cache[Key{s1, s2, s3}]
	}

	result := Value{false, s1_counts, s2_counts}
	// check start of s1 and s2
	if s1[0] == s3[0] {
		// fmt.Println("s1[0] == s3[current_index]")
		// dfs on s1
		return_value := dfs(s1[1:], s2, s3[1:], s1_counts, s2_counts, cache)
		if return_value.result {
			result = return_value
		}
	} else {
		result.s1_counts += 1
	}
	if s2[0] == s3[0] {
		// fmt.Println("s2[0] == s3[current_index]")
		// dfs on s2
		return_value := dfs(s1, s2[1:], s3[1:], s1_counts, s2_counts, cache)
		if return_value.result {
			result = return_value
		}
	} else {
		result.s2_counts += 1
	}

	cache[Key{s1, s2, s3}] = result

	return result

}

type Key struct {
	s1 string
	s2 string
	s3 string
}
type Value struct {
	result    bool
	s1_counts int
	s2_counts int
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	cache := map[Key]Value{}

	return dfs(s1, s2, s3, 0, 0, cache).result
}
