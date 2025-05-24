package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	src    int
	dist   int
	weight int
}

type Graph struct {
	vertices int
	edegs    []Edge
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edegs:    make([]Edge, 0),
	}
}

func (g *Graph) AddEdge(src, dst, weight int) {
	g.edegs = append(g.edegs, Edge{src: src, dist: dst, weight: weight})
}

func (g *Graph) KruskalMST() (int, [][]int) {

	sort.Slice(g.edegs, func(i, j int) bool {
		return g.edegs[i].weight < g.edegs[j].weight
	})

	parent := make([]int, g.vertices)
	rank := make([]int, g.vertices)

	for i := 0; i < g.vertices; i++ {
		parent[i] = i
	}

	mstEdge := make([][]int, 0)
	totalWeight := 0
	edgeCount := 0

	for i := 0; edgeCount < g.vertices-1 && i < len(g.edegs); i++ {
		nexrEdge := g.edegs[i]

	}

}

func (g *Graph) find(parent []int, i int) int {

	if parent[i] != i {
		parent[i] = g.find(parent, parent[i])
	}

	return parent[i]
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

	// totalWeight, mstEdges := graph.KruskalMST()

	// fmt.Println("Minimum Spanning Tree Edges (Kruskal's):")
	// for _, edge := range mstEdges {
	// 	fmt.Printf("%d - %d (weight: %d)\n", edge[0], edge[1], edge[2])
	// }
	// fmt.Printf("Total weight: %d\n", totalWeight)

	fmt.Println(graph)
}
