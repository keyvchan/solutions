package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solution("...XXX..X....XXX", 7))
}

type Segement struct {
	start int
	end   int
	value int
}

type Segements []Segement

func (s Segements) Len() int {
	return len(s)
}

func (s Segements) Less(i, j int) bool {
	return s[i].value < s[j].value
}

func (s Segements) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Solution(S string, B int) int {
	// write your code in Go (Go 1.4)
	pathholes := []Segement{}

	seg := Segement{}
	for i := 0; i < len(S); i++ {
		if S[i] == '.' {
			// check if seg is empty
			if seg.start != seg.end {
				// put it into map
				seg.value = seg.end - seg.start
				pathholes = append(pathholes, seg)
			}
			seg = Segement{}
			continue
		}
		if S[i] == 'X' {
			// check if we are in a segement
			if seg.start == seg.end {
				// create a Segement
				seg = Segement{start: i, end: i + 1}
			} else {
				// put current segement in the map
				seg.end = i + 1
			}
		}

	}
	// if segement is not empty put it in the map
	if seg.start != seg.end {
		seg.value = seg.end - seg.start
		pathholes = append(pathholes, seg)
	}

	sort.Sort(Segements(pathholes))

	fixed := 0
	// fix
	for i := 0; i < len(pathholes); i++ {
		if pathholes[i].value > B {
			// fix as many as possible
			fixed += B
			break
		} else {
			B -= pathholes[i].value + 1
			fixed += pathholes[i].value
		}
	}

	return fixed
}
