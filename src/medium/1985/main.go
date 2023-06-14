package main

import (
	"fmt"
	"math/big"
	"sort"
)

func main() {
	fmt.Println(kthLargestNumber([]string{"3", "2", "1", "5", "6", "4"}, 2))
	fmt.Println(kthLargestNumber([]string{"683339452288515879", "7846081062003424420", "4805719838", "4840666580043", "83598933472122816064", "522940572025909479", "615832818268861533", "65439878015", "499305616484085", "97704358112880133", "23861207501102", "919346676", "60618091901581", "5914766072", "426842450882100996", "914353682223943129", "97", "241413975523149135", "8594929955620533", "55257775478129", "528", "5110809", "7930848872563942788", "758", "4", "38272299275037314530", "9567700", "28449892665", "2846386557790827231", "53222591365177739", "703029", "3280920242869904137", "87236929298425799136", "3103886291279"},
		3))
	fmt.Println(kthLargestNumber([]string{"0", "0"}, 2))
}

func kthLargestNumber(nums []string, k int) string {
	// sort the array
	sort.Slice(nums, func(i, j int) bool {
		bi := big.NewInt(0)
		bi.SetString(nums[i], 10)

		bj := big.NewInt(0)
		bj.SetString(nums[j], 10)

		if bi.Cmp(bj) == -1 {
			return true
		} else {
			return false
		}
	})

	fmt.Println(nums)
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }
	return nums[len(nums)-k]
}
