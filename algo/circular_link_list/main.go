package main

import "math/rand"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

func New() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (ll *CircularLinkedList) Append(data int) {

	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}
	currentNode := ll.head
	for currentNode.next != nil {

		currentNode = currentNode.next
	}
	currentNode.next = newNode
}

func (ll *CircularLinkedList) MakeItCircularAt(index int) {

	count := 0

	current := ll.head
	var lastNode *Node
	var toNode *Node

	for current != nil {

		lastNode = current
		if count == index {
			toNode = current
		}
		count++
		current = current.next
	}

	lastNode.next = toNode

}
func main() {

	cc := New()

	for i := 0; i < 10; i++ {
		cc.Append(rand.Intn(100))
	}

	cc.MakeItCircularAt(5)

}
