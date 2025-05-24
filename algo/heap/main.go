package main

import (
	"fmt"
	"math/rand"
)

type Arr struct {
	Array []int
}

type Heap struct {
	heap []int
}

func NewArray() *Arr {
	return &Arr{}
}

func NewHeap() *Heap {
	return &Heap{}
}

func (arr *Arr) Insert() {

	for i := 1; i < 10; i++ {
		arr.Array = append(arr.Array, rand.Intn(100))
	}
}

func (heap *Heap) InsertIntoHeap(value int) {

	heap.heap = append(heap.heap, value)
	node := len(heap.heap) - 1
	if node == 0 {
		return
	}

	heap.RearrangeHeap(node, (node-1)/2)
}

func (heap *Heap) RearrangeHeap(node int, parantNode int) {

	if parantNode < 0 {
		return
	} else if heap.heap[parantNode] < heap.heap[node] {
		heap.heap[parantNode], heap.heap[node] = heap.heap[node], heap.heap[parantNode]
		heap.RearrangeHeap(parantNode, (parantNode-1)/2)
	}

}

func (heap *Heap) ReArrangeHeapDelete(pareantNode int) {

	leftChildNode := (2 * pareantNode) + 1
	rightChildNode := (2 * pareantNode) + 2
	largerNode := pareantNode

	if leftChildNode < len(heap.heap) && heap.heap[largerNode] < heap.heap[leftChildNode] {

		largerNode = leftChildNode

	}
	if rightChildNode < len(heap.heap) && heap.heap[largerNode] < heap.heap[rightChildNode] {

		largerNode = rightChildNode

	}

	if pareantNode != largerNode {
		heap.heap[largerNode], heap.heap[pareantNode] = heap.heap[pareantNode], heap.heap[largerNode]
		heap.ReArrangeHeapDelete(largerNode)
	}

}

func (heap *Heap) Delete() int {

	out := heap.heap[0]

	heap.heap[0] = heap.heap[len(heap.heap)-1]
	heap.heap = heap.heap[:len(heap.heap)-1]

	heap.ReArrangeHeapDelete(0)

	return out
}

func (heap *Heap) HeapSort(arr []int) []int {

	var sortedArray []int

	for _, val := range arr {
		heap.InsertIntoHeap(val)
	}

	for _, _ = range heap.heap {
		sortedArray = append(sortedArray, heap.Delete())
	}
	return sortedArray
}

func (heap *Heap) HeapFy(array []int, len int, index int) []int {


	leftChildNode := 2*index+1
	rightChildNode:= 2*index+2
	largeNimber := index


	if leftChildNode< len && array[f]


}

// func (heap *Heap)IsCompleteHeap(arr []int )bool{

// if len

// }

func main() {

	arr := NewArray()
	arr.Insert()

	heap := NewHeap()

	// for _, val := range arr.Array {
	// 	heap.InsertIntoHeap(val)
	// 	fmt.Println(heap.heap)
	// }

	// for _, _ = range heap.heap {
	// 	fmt.Println(heap.Delete())
	// 	fmt.Println(heap.heap)
	// }

	for i := 0; i <= (len(heap.heap)/(2-1))-1-1; i++ {

		heap.HeapFy(arr.Array, len(heap.heap), i)

	}

	fmt.Println(arr.Array)
	fmt.Println(heap.HeapSort(arr.Array))

}
