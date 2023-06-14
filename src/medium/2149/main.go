package main

func main() {

}

func rearrangeArray(nums []int) []int {

	// take out all the
	postives := []int{}
	negatives := []int{}
	for _, num := range nums {
		if num < 0 {
			negatives = append(negatives, num)
		} else {
			postives = append(postives, num)
		}
	}
	j := 0
	for i := 0; i < len(postives); i++ {
		nums[j] = postives[i]
		nums[j+1] = negatives[i]
		j += 2
	}

	return nums
}
