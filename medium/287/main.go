package main

func main() {

}

func findDuplicate(nums []int) int {
	slow_pointer := nums[0]
	faster_pointer := nums[0]
	another_slow_pointer := nums[0]

	for {
		// slow_pointer steps forward 1 step
		slow_pointer = nums[slow_pointer]
		// faster_pointer steps forward 2 steps
		faster_pointer = nums[nums[faster_pointer]]
		if slow_pointer == faster_pointer {
			// intersection point
			// set another slow_pointer to the start of the array
			for {
				// slow_pointer steps forward 1 step
				slow_pointer = nums[slow_pointer]
				// another_slow_pointer steps forward 1 step
				another_slow_pointer = nums[another_slow_pointer]
				if slow_pointer == another_slow_pointer {
					return slow_pointer
				}
			}
		}
	}

}
