package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	pointer1 := head
	pointer2 := head
	if pointer1 == nil {
		return false
	}

	for pointer2.Next != nil {
		// pointer1 move one step
		pointer1 = pointer1.Next

		// pointer2 move two steps
		pointer2 = pointer2.Next.Next
		if pointer1 == pointer2 {
			return true
		}
		if pointer2 == nil {
			return false
		}
	}
	return false
}
