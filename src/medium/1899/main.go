package main

func main() {

}

func mergeTriplets(triplets [][]int, target []int) bool {
	elements := 0
	equal_flag_0 := false
	equal_flag_1 := false
	equal_flag_2 := false

	for _, triplets := range triplets {
		if triplets[0] <= target[0] && triplets[1] <= target[1] && triplets[2] <= target[2] {
			if triplets[0] == target[0] {
				equal_flag_0 = true
			}
			if triplets[1] == target[1] {
				equal_flag_1 = true
			}
			if triplets[2] == target[2] {
				equal_flag_2 = true
			}
			// add current triplets to elements
			elements += 1
		}
	}
	if elements == 0 {
		return false
	}
	if equal_flag_0 && equal_flag_1 && equal_flag_2 {
		return true
	} else {
		return false
	}

}
