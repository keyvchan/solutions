package main

func main() {

}

type NumArray struct {
	nums               []int
	range_sum          map[int]int
	last_changed_index int
}

func Constructor(nums []int) NumArray {
	// rage sum
	range_sum := map[int]int{}
	for i := 0; i < len(nums); i++ {
		range_sum[i] = nums[i] + range_sum[i-1]
	}
	return NumArray{nums: nums, range_sum: range_sum, last_changed_index: len(nums)}

}

func (this *NumArray) Update(index int, val int) {
	// update the value of index
	this.nums[index] = val
	// update the range sum
	this.last_changed_index = index

}

func (this *NumArray) SumRange(left int, right int) int {
	// if left greater than last changed index, recalculate the range sum
	if right < this.last_changed_index {
		return this.range_sum[right] - this.range_sum[left-1]
	} else {
		// recalculate the range sum
		for i := this.last_changed_index; i <= right; i++ {
			this.range_sum[i] = this.nums[i] + this.range_sum[i-1]
		}
		this.last_changed_index = right
		return this.range_sum[right] - this.range_sum[left-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
