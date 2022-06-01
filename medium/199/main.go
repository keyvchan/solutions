package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	// bfs
	res := []int{}
	queue := list.New()
	queue.PushBack(root)

	for {
		queue_len := queue.Len()
		// take node from queue
		for queue_len > 0 {
			elememt := queue.Remove(queue.Front())
			if elememt != nil {
				queue.PushBack(elememt.(*TreeNode).Left)
			}
			if elememt != nil {
				queue.PushBack(elememt.(*TreeNode).Right)
			}
			queue_len -= 1
		}
		// back of queue is the rightest node
		res = append(res, queue.Back().Value.(*TreeNode).Val)
		if queue_len == 0 {
			break
		}
	}

	return res
}
