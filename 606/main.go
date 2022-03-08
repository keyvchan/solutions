package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(tree2str(tree1))
}

func pre_order(root *TreeNode, str *string) {
	if root == nil {
		return
	}
	*str += "("
	*str += fmt.Sprint(root.Val)
	if root.Left == nil && root.Right != nil {
		*str += "()"
	}
	pre_order(root.Left, str)
	pre_order(root.Right, str)
	*str += ")"
}

func tree2str(root *TreeNode) string {
	str := ""

	pre_order(root, &str)

	return str[1 : len(str)-1]
}
