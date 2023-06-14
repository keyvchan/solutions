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
	if root == nil {
		return []int{}
	}
	// bfs
	res := []int{root.Val}
	queue := list.New()
	queue.PushBack(root)

	for {
		queue_len := queue.Len()
		// take node from queue
		for queue_len > 0 {
			elememt := queue.Remove(queue.Front())
			if elememt.(*TreeNode) != nil {
				if elememt.(*TreeNode).Left != nil {
					queue.PushBack(elememt.(*TreeNode).Left)
				}
				if elememt.(*TreeNode).Right != nil {
					queue.PushBack(elememt.(*TreeNode).Right)
				}
			}
			queue_len -= 1
		}
		if queue.Len() == 0 {
			break
		}
		// back of queue is the rightest node
		res = append(res, queue.Back().Value.(*TreeNode).Val)
	}

	return res
}
