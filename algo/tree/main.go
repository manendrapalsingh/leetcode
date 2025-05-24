package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func NewTree() *BST {
	return &BST{}
}

func (t *BST) Insert(data int) {
	t.root = t.InsertRec(t.root, data)
}

func (t *BST) InsertRec(node *Node, val int) *Node {

	if node == nil {
		return &Node{data: val}
	}

	if val < node.data {

		node.left = t.InsertRec(node.left, val)
	}
	if val > node.data {
		node.right = t.InsertRec(node.right, val)
	}

	return node

}

func (t *BST) PreOrder() {
	PreOrderRec(t.root)
	fmt.Println()
}

func PreOrderRec(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.data)
	PreOrderRec(node.left)
	PreOrderRec(node.right)
}

func (t *BST) Search(data int) bool {
	return SearchRec(t.root, data)
}

func SearchRec(node *Node, data int) bool {
	output := false

	if node == nil {
		return output
	}

	if data == node.data {
		return true
	}

	if data < node.data {
		output = SearchRec(node.left, data)
	} else {
		output = SearchRec(node.right, data)
	}

	return output
}

func (t *BST) Delete(data int) {
	t.root = t.deleteRec(t.root, data)
}

func (t *BST) deleteRec(node *Node, data int) *Node {

	if node == nil {
		return nil
	}

	if data < node.data {
		node.left = t.deleteRec(node.left, data)
	} else if data > node.data {
		node.right = t.deleteRec(node.right, data)
	} else {

		if node.left == nil && node.right == nil {
			return nil
		}

		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		successor := t.findMin(node.right)
		node.data = successor.data
		node.right = t.deleteRec(node.right, data)
	}

	return node
}

func (t *BST) Height() int {
	var x, y int

}
func (t *BST) findMin(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}

	return current
}
func (t *BST) InOrder() {

	InOrderRec(t.root)
}

func InOrderRec(node *Node) {

	if node == nil {
		return
	}

	if node.left != nil {
		InOrderRec(node.left)
	}
	fmt.Printf("%v ", node.data)

	if node.right != nil {
		InOrderRec(node.right)
	}
}

func main() {

	t1 := NewTree()

	for i := 0; i < 10; i++ {
		t1.Insert(rand.Intn(1000))
	}

	t1.PreOrder()

	fmt.Println(t1.Search(15))

	t1.InOrder()
}
