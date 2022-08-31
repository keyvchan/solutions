package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var last *ListNode
	var head *ListNode
	for {

		var min *ListNode
		for i := 0; i < len(lists); i++ {
			if min == nil || (lists[i] != nil && lists[i].Val < min.Val) {
				min = lists[i]
			}
		}
		if min == nil {
			break
		}

		if last == nil {
			last = min
			head = last
		} else {
			last.Next = min
			last = last.Next
		}

		// search for the current min
		for i := 0; i < len(lists); i++ {
			if lists[i] == min {
				lists[i] = lists[i].Next
				break
			}
		}

	}

	return head
}
