package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node   int
	weight int
}

type Graph struct {
	vertices int
	adjList  map[int][]Edge
}

type HeapItem struct {
	vertex int
	weight int
	parent int
}

type Heap []*HeapItem

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*HeapItem))
}
func (h *Heap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]Edge),
	}
}

func (g *Graph) AddEdge(src, dest, weight int) {
	g.adjList[src] = append(g.adjList[src], Edge{node: dest, weight: weight})
	g.adjList[dest] = append(g.adjList[dest], Edge{node: src, weight: weight})
}

func (g *Graph) PrimMST(startVertex int) (int, [][]int) {

	visited := make([]bool, g.vertices)

	minHeap := &Heap{}
	heap.Init(minHeap)

	mstEdges := make([][]int, 0)
	totalWeight := 0

	heap.Push(minHeap, &HeapItem{vertex: startVertex, weight: 0, parent: -1})

	for len(*minHeap) > 0 {

		curent := heap.Pop(minHeap).(*HeapItem)

		if visited[curent.vertex] {
			continue
		}

		visited[curent.vertex] = true
		totalWeight += curent.weight

		if curent.parent != -1 {
			mstEdges = append(mstEdges, []int{curent.parent, curent.vertex, curent.weight})
		}

		for _, edge := range g.adjList[curent.vertex] {
			if !visited[edge.node] {
				heap.Push(minHeap, &HeapItem{
					vertex: edge.node,
					weight: edge.weight,
					parent: curent.parent,
				})
			}
		}

	}

	return totalWeight, mstEdges
}

func main() {
	graph := NewGraph(5)
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 3, 6)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(1, 3, 8)
	graph.AddEdge(1, 4, 5)
	graph.AddEdge(2, 4, 7)
	graph.AddEdge(3, 4, 9)

	totalWeight, mstEdges := graph.PrimMST(0)

	fmt.Println("Minimum Spanning Tree Edges:")
	for _, edge := range mstEdges {
		fmt.Printf("%d - %d (weight: %d)\n", edge[0], edge[1], edge[2])
	}
	fmt.Printf("Total weight: %d\n", totalWeight)
}
