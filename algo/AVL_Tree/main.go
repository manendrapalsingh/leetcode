package main

import "fmt"

type Node struct {
	data   int
	height int
	left   *Node
	right  *Node
}

type AVLTree struct {
	root *Node
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Insert(data int) {
	t.root = t.InsertRec(t.root, data)
}

func (t *AVLTree) InsertRec(node *Node, data int) *Node {
	if node == nil {
		return &Node{data: data, height: 1}
	}

	if data < node.data {
		node.left = t.InsertRec(node.left, data)
	} else {
		node.right = t.InsertRec(node.right, data)
	}

	// Update height
	node.height = max(height(node.left), height(node.right)) + 1

	// Balance the node
	balance := balanceFactor(node)

	// LL Case
	if balance > 1 && data < node.left.data {
		return t.LLRotation(node)
	}

	// RR Case
	if balance < -1 && data > node.right.data {
		return t.RRRotation(node)
	}

	// LR Case
	if balance > 1 && data > node.left.data {
		return t.LRRotation(node)
	}

	// RL Case
	if balance < -1 && data < node.right.data {
		return t.RLRotation(node)
	}

	return node
}

// ----------- ROTATIONS -----------

func (t *AVLTree) LLRotation(node *Node) *Node {
	leftChild := node.left
	node.left = leftChild.right
	leftChild.right = node

	// Update heights
	node.height = max(height(node.left), height(node.right)) + 1
	leftChild.height = max(height(leftChild.left), height(leftChild.right)) + 1

	return leftChild
}

func (t *AVLTree) RRRotation(node *Node) *Node {
	rightChild := node.right
	node.right = rightChild.left
	rightChild.left = node

	// Update heights
	node.height = max(height(node.left), height(node.right)) + 1
	rightChild.height = max(height(rightChild.left), height(rightChild.right)) + 1

	return rightChild
}

func (t *AVLTree) LRRotation(node *Node) *Node {
	node.left = t.RRRotation(node.left)
	return t.LLRotation(node)
}

func (t *AVLTree) RLRotation(node *Node) *Node {
	node.right = t.LLRotation(node.right)
	return t.RRRotation(node)
}

// ----------- UTILITIES -----------

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func balanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.left) - height(node.right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ----------- TRAVERSAL -----------

func (t *AVLTree) InOrder() {
	fmt.Print("InOrder: ")
	InOrderRec(t.root)
	fmt.Println()
}

func InOrderRec(node *Node) {
	if node == nil {
		return
	}
	InOrderRec(node.left)
	fmt.Printf("->%v", node.data)
	InOrderRec(node.right)
}

// ----------- MAIN -----------

func main() {
	t := NewAVLTree()

	t.Insert(10)
	t.Insert(5)
	t.Insert(2) // LL
	t.Insert(8) // LR
	t.Insert(15)
	t.Insert(20) // RR
	t.Insert(18) // RL

	t.InOrder()
	fmt.Println("Root:", t.root.data)
}
