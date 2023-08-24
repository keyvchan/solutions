package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumUnits([][]int{{2, 1}, {2, 3}, {4, 5}}, 4))
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {

	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	units := 0
	for _, box := range boxTypes {
		for i := 0; i < box[0]; i++ {
			if truckSize-1 >= 0 {
				truckSize--
				units += box[1]
			} else {
				break
			}
		}
		if truckSize == 0 {
			break
		}
	}
	return units

}
