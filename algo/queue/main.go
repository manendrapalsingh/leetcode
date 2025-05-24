package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	data int
	next *Node
}

type LinkedQueue struct {
	head *Node
	tail *Node
}
type Queue struct {
	data  []int
	last  int
	front int
	size  int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func NewQueue(size int) *Queue {
	return &Queue{
		data:  make([]int, size),
		last:  0,
		size:  size,
		front: -1,
	}
}

func NewCircularQueue(size int) *Queue {
	return &Queue{
		data:  make([]int, size),
		last:  0,
		size:  size,
		front: 0,
	}
}
func (q *Queue) Enqueue(v int) {

	if q.last < q.size {
		q.data[q.last] = v
		q.last++
	} else {
		q.data = append(q.data, v)
	}
}

func (q *Queue) Dequeue() int {

	var out int
	if q.last == q.front {
		fmt.Errorf("q is empty")
		return out
	}

	out = q.data[q.front+1]
	q.front = q.front + 1

	return out
}

func (q *Queue) CircularEnqueue(val int) {

	if (q.last+1)%q.size == q.front {
		fmt.Errorf("q is is full")
		return
	} else {
		q.last = (q.last + 1) % q.size
		q.data[q.last] = val
		return
	}

}

func (q *Queue) CircularDequeue() int {
	var out int

	if q.last == q.front {
		fmt.Errorf("q is empty")
	} else {

		q.front = (q.front + 1) % q.size
		out = q.data[q.front]
	}

	return out
}

func (q *LinkedQueue) Enqueue(v int) {

	newNode := &Node{data: v}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
}

func (q *LinkedQueue) Print() {

	if q.head == nil {
		return
	}

	current := q.head

	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}

	current2 := q.tail

	fmt.Println()

	for current2 != nil {
		fmt.Printf("tail =%d", current2.data)

		current2 = current2.next
	}

}

func (q *LinkedQueue) Dequeue() int {

	var out int

	if q.head == nil {
		fmt.Errorf("q is empty")
	}
	current := q.head
	out = current.data
	q.head = current.next

	if q.head == nil {
		q.tail = nil
	}
	return out
}
func main() {

	q1 := NewQueue(10)

	for i := 0; i < 10; i++ {
		q1.Enqueue(rand.Intn(100))
	}

	q2 := NewCircularQueue(10)
	for i := 0; i < 10; i++ {
		q2.CircularEnqueue(rand.Intn(100))
	}

	fmt.Println(q2.data)

	fmt.Println(q2.CircularDequeue())
	q2.CircularEnqueue(rand.Intn(100))
	fmt.Println(q2.data)
	fmt.Println(q2.CircularDequeue())

	q3 := NewLinkedQueue()

	for i := 0; i < 10; i++ {
		q3.Enqueue(rand.Intn(100))
	}

	q3.Print()
	fmt.Println("ksdsdjksjd")
	fmt.Println(q3.Dequeue())
	fmt.Println("dnasdjjsad")
	fmt.Println(q3.Dequeue())
	q3.Print()

}
