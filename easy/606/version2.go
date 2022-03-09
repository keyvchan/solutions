package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(tree2str(tree1))
}

func pre_order(root *TreeNode, str strings.Builder) {
	if root == nil {
		return
	}
	str.Write([]byte("("))
	str.Write([]byte(fmt.Sprint(root.Val)))
	if root.Left == nil && root.Right != nil {
		str.Write([]byte("()"))
	}
	pre_order(root.Left, str)
	pre_order(root.Right, str)
	str.Write([]byte(")"))
}

func tree2str(root *TreeNode) string {
	strs := strings.Builder{}
	strs.Grow(100)

	strs.Write([]byte(fmt.Sprint(root.Val)))
	pre_order(root.Left, strs)
	pre_order(root.Right, strs)

	return strs.String()
}
