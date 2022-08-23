package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// put all nodes in a array
	node_values := []int{}
	for head != nil {
		node_values = append(node_values, head.Val)
		head = head.Next
	}

	for i, j := 0, len(node_values)-1; ; {
		if i > j {
			return true
		}
		if node_values[i] != node_values[j] {
			return false
		}
		i++
		j--

	}

}
