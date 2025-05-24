package main

import (
	"fmt"
	"math/rand"
)

type Stack struct {
	data []int
	size int
	top  int
}

type Node struct {
	data     int
	previous *Node
}

type Stack2 struct {
	top *Node
}

func newStack2() *Stack2 {
	return &Stack2{}
}

func (s *Stack2) Push(data int) {

	newNode := &Node{data: data}

	if s.top == nil {
		s.top = newNode
		return
	} else {
		newNode.previous = s.top
		s.top = newNode
	}

}

func (s *Stack2) Pop() {

	if s.top == nil {
		return
	}

	s.top = s.top.previous
}

func (s *Stack2) Print() {

	if s.top == nil {
		return
	}

	current := s.top

	for current != nil {
		fmt.Printf("%v->", current.data)
		current = current.previous
	}
}
func NewStack(size int) *Stack {

	return &Stack{
		data: make([]int, size),
		top:  -1,
		size: size,
	}
}

func (s *Stack) Push(v int) {

	if s.top+1 < s.size {
		index := s.top + 1
		s.data[index] = v
		s.top = index
	} else {
		s.data = append(s.data, v)
		s.top++
	}

}

func (s *Stack) Pop() {

	if s.top == -1 {
		return
	}

	s.data = s.data[0:s.top]
	s.top = s.top - 1

}

func (s *Stack) Peak() int {
	return s.data[s.top]
}

func main() {

	stack := NewStack(10)

	for i := 0; i < 10; i++ {
		stack.Push(rand.Intn(100))
	}

	fmt.Println(stack.data)
	stack.Push(999)
	fmt.Println(stack.data)

	stack2 := newStack2()

	for i := 0; i < 10; i++ {
		stack2.Push(i)
	}

	stack2.Print()

	fmt.Println()

	stack2.Pop()
	stack2.Print()

}
