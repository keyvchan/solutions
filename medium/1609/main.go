package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	// bfs
	queue := list.New()

	next_level := list.New()

	queue.PushBack(root)

	count := 0
	for queue.Len() > 0 {

		last := 0
		for queue.Len() > 0 {
			element := queue.Remove(queue.Front()).(*TreeNode)
			// verify
			if count%2 == 0 {
				if element.Val%2 == 0 {
					return false
				}
				if last != 0 {
					// check if its greater than last
					if element.Val <= last {
						return false
					}
				}
			}
			if count%2 == 1 {
				if element.Val%2 == 1 {
					return false
				}
				if last != 0 {
					// check if its less than last
					if element.Val >= last {
						return false
					}
				}
			}
			last = element.Val

			if element.Left != nil {
				next_level.PushBack(element.Left)
			}
			if element.Right != nil {
				next_level.PushBack(element.Right)
			}
		}
		queue, next_level = next_level, queue
		count++
	}

	return true
}
