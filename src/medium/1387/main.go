package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getKth(12, 15, 2))
	fmt.Println(getKth(7, 11, 4))
}

type Pair struct {
	i     int
	power int
}

func getKth(lo int, hi int, k int) int {

	power_array := make([]Pair, 0, hi-lo+1)

	for i := lo; i <= hi; i++ {
		// calculate power of i
		power := 0
		temp := i
		for temp != 1 {
			if temp%2 == 0 {
				// even
				temp = temp / 2
			} else {
				// odd
				temp = temp*3 + 1
			}
			power += 1
		}
		power_array = append(power_array, Pair{i, power})
	}

	sort.Slice(power_array, func(i, j int) bool {
		if power_array[i].power == power_array[j].power {
			// power equal, sort by i
			return power_array[i].i < power_array[j].i

		} else {
			return power_array[i].power < power_array[j].power
		}
	})

	kth := power_array[k-1]
	fmt.Println(power_array)
	return kth.i

}
