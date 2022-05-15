package main

func main() {

}

type NumArray struct {
	array []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	new_array := this.array[left : right+1]
	sum := 0
	for _, v := range new_array {
		sum += v
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
