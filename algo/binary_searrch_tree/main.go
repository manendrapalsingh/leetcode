package main

type TreeNode struct {
	key   int
	val   int
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func (bt *BinaryTree) Insert(key int, val int) {

	n := &TreeNode{key: key, val: val}
	if bt.root == nil {
		bt.root = n
	} else {
		insertNode
	}
}

func (bt *BinaryTree) insertNode(TreeNode, newNode *TreeNode) {
	
}
