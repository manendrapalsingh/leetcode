package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data int
	next *Node
}

type LinkList struct {
	head *Node
}

func InitLinkList() *LinkList {
	return &LinkList{}
}

func (ll *LinkList) Append(val int) {

	newNode := &Node{val, nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}

func Length(ll *LinkList) int {

	count := 0
	for ll.head == nil {
		return count
	}
	current := ll.head

	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (ll *LinkList) InsertAt(val, index int) {

	length := Length(ll)
	newNode := &Node{val, nil}

	if length == 0 && index == 0 {
		ll.head = newNode
		return
	} else if index > length {
		fmt.Errorf("index out of range")
		return
	}

	current := ll.head

	for i := 0; i < length; i++ {

		if i == index-1 {
			newNode.next = current.next
			current.next = newNode

			return
		}

		current = current.next
	}

}

func (ll *LinkList) Traverse() {

	sum := 0
	count := 0
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%v->", current.data)
		sum = sum + current.data
		count++
		current = current.next

	}

	println()

}

func (ll *LinkList) Search(val int) bool {

	if ll.head == nil {
		return false
	}

	current := ll.head

	for current != nil {
		if current.data == val {
			return true
		}
		current = current.next
	}

	return false

}

func (ll *LinkList) InsertBefore(val int) {

	newNode := &Node{val, nil}

	if ll.head == nil {
		ll.head = newNode
		return
	} else {
		newNode.next = ll.head
		ll.head = newNode
		return
	}
}

func (ll *LinkList) SortedInsert(val int) {
	newNode := &Node{val, nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	var prev *Node

	for current != nil {

		if val < current.data && prev != nil {

			newNode.next = prev.next
			prev.next = newNode
			return

		} else if val < current.data && prev == nil {
			newNode.next = current
			ll.head = newNode
			return
		} else if val > current.data && current.next == nil {

			current.next = newNode
		}
		prev = current
		current = current.next

	}

}

func (ll *LinkList) Delete(index int) {

	if ll.head == nil {
		return
	}

	count := 0
	current := ll.head

	if index == 0 {
		ll.head = current.next
		return
	}

	var prev *Node

	for current != nil {
		if count == index {
			prev.next = current.next
			return
		}

		count++
		prev = current
		current = current.next

	}
}

func (ll *LinkList) IsSorted() bool {

	current := ll.head

	prev := current
	for current != nil {

		if current.next == nil {
			return true
		}

		if (current.data < prev.data || current.data > current.next.data) && current != prev {
			return false
		}

		prev = current
		current = current.next

	}

	return false
}

func (ll *LinkList) RemoveDuplicates() {

	if ll.head == nil {
		return
	}

	current := ll.head

	var perv *Node

	for current.next != nil {

		if perv == nil {

			perv = current
			current = current.next
			continue
		} else if perv.data == current.data {
			perv.next = current.next
			current = current.next
			continue

		}

		perv = current
		current = current.next

	}

}

func (ll *LinkList) Reverse() {

	if ll.head == nil {
		return
	}

	current := ll.head

	var prev *Node
	var prev_prev *Node

	for current != nil {

		prev_prev = prev
		prev = current
		current = current.next

		prev.next = prev_prev

	}

	ll.head = prev
}

func Concatenate(ll1 *LinkList, ll2 *LinkList) {

	if ll1.head == nil {
		ll1.head = ll2.head
		return
	}

	current := ll1.head

	for current.next != nil {
		current = current.next
	}

	current.next = ll2.head

}

func Merge(ll1 *LinkList, ll2 *LinkList) {

	if ll1.head == nil {
		ll1.head = ll2.head
	}

	current1 := ll1.head
	current2 := ll2.head

	prev1 := current1
	prev2 := current2

	for current1 != nil && current2 != nil {

		if current1.data <= current2.data {
			prev1 = current1
			current1 = current1.next
		} else {

			current2.next = current1

		}

	}
}
func main() {

	//linkList1 := InitLinkList()
	//for i := 0; i < 10; i++ {
	//	linkList1.Append(i)
	//}

	//fmt.Println(linkList1.Search(20))
	//fmt.Println(linkList1.Search(5))
	//
	//fmt.Println(Length(linkList1))

	//linkList1.InsertAt(99, 2)
	//linkList1.InsertBefore(99)
	//linkList1.SortedInsert(30)
	//linkList1.Delete(0)
	//linkList1.Delete(4)
	//linkList1.Delete(9)

	//linkList1.InsertAt(3, 3)
	//linkList1.InsertAt(3, 4)

	//linkList1.Traverse()

	//linkList1.RemoveDuplicates()

	//linkList1.Reverse()

	//linkList2 := InitLinkList()
	//
	//for i := 10; i < 20; i++ {
	//	linkList2.Append(i)
	//}
	//
	//Concatenate(linkList1, linkList2)
	//linkList1.Traverse()

	ll1 := InitLinkList()
	ll2 := InitLinkList()

	for i := 0; i < 10; i++ {
		ll1.SortedInsert(rand.Intn(100))
		ll2.SortedInsert(rand.Intn(100))
	}

	ll1.Traverse()
	ll2.Traverse()

	Merge(ll1, ll2)

	ll1.Traverse()

}
