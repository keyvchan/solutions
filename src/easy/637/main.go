package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	result := []float64{}

	next_level := list.New()

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		sum := 0
		size := queue.Len()
		for queue.Len() > 0 {
			elememt := queue.Remove(queue.Front()).(*TreeNode)

			if elememt.Left != nil {
				next_level.PushBack(elememt.Left)
			}
			if elememt.Right != nil {
				next_level.PushBack(elememt.Right)
			}
			sum += elememt.Val
		}

		result = append(result, float64(sum)/float64(size))
		queue.PushBackList(next_level)
		next_level.Init()
	}

	return result
}
