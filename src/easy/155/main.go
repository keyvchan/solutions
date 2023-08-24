package main

func main() {

}

type MinStack struct {
	array []int
	min   []int
}

func Constructor() MinStack {
	array := MinStack{
		array: []int{},
		min:   []int{0},
	}
	return array
}

func (this *MinStack) Push(val int) {
	if len(this.min) == 0 {
		this.min = append(this.min, val)
	} else {
		if this.min[len(this.min)-1] >= val {
			this.min = append(this.min, val)
		}
	}

	this.array = append(this.array, val)
}

func (this *MinStack) Pop() {
	if this.min[len(this.min)-1] == this.array[len(this.array)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.array = this.array[:len(this.array)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.array) == 0 {
		return 0
	}
	return this.array[len(this.array)-1]
}
