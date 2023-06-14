package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Root   *TreeNode
	Cursor int
	Arr    []*TreeNode
}

func dfs(root *TreeNode, array *[]*TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left, array)
	*array = append(*array, root)
	dfs(root.Right, array)
}

func Constructor(root *TreeNode) BSTIterator {
	// binary tree flatten
	array := []*TreeNode{}
	dfs(root, &array)
	// the first element should be the smallest element
	array = append([]*TreeNode{
		&TreeNode{
			Val:   array[0].Val - 1,
			Left:  nil,
			Right: nil,
		},
	}, array...)

	return BSTIterator{
		Root:   root,
		Cursor: 0,
		Arr:    array,
	}
}

func (this *BSTIterator) Next() int {
	this.Cursor++
	return this.Arr[this.Cursor].Val
}

func (this *BSTIterator) HasNext() bool {

	if this.Cursor < len(this.Arr)-1 {
		return true
	} else {
		return false
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
