package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
    head = head.Next
	}

	// swap from back to front
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}

	return result

}
