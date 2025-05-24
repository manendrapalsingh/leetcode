package main

import (
	"container/heap"
	"fmt"
	"strings"
)

// Node represents a node in the Huffman tree
type Node struct {
	char      rune
	frequency int
	left      *Node
	right     *Node
}

// Heap implements heap.Interface for Node
type Heap []*Node

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// buildFrequencyMap creates a frequency map for each character in the input string
func buildFrequencyMap(data string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range data {
		freq[char]++
	}
	return freq
}

// buildHuffmanTree constructs the Huffman tree from the frequency map
func buildHuffmanTree(freq map[rune]int) *Node {
	h := &Heap{}
	heap.Init(h)

	// Create a leaf node for each character and add to the heap
	for char, count := range freq {
		heap.Push(h, &Node{char: char, frequency: count})
	}

	// Build the tree by combining nodes until only one remains
	for h.Len() > 1 {
		left := heap.Pop(h).(*Node)
		right := heap.Pop(h).(*Node)

		parent := &Node{
			frequency: left.frequency + right.frequency,
			left:      left,
			right:     right,
		}

		heap.Push(h, parent)
	}

	return heap.Pop(h).(*Node)
}

// buildCodeMap generates the Huffman codes by traversing the tree
func buildCodeMap(root *Node, code string, codeMap map[rune]string) {
	if root == nil {
		return
	}

	// Leaf node contains a character
	if root.char != 0 {
		codeMap[root.char] = code
		return
	}

	buildCodeMap(root.left, code+"0", codeMap)
	buildCodeMap(root.right, code+"1", codeMap)
}

// encode converts the input string to its Huffman encoded binary string
func encode(data string, codeMap map[rune]string) string {
	var builder strings.Builder
	for _, char := range data {
		builder.WriteString(codeMap[char])
	}
	return builder.String()
}

// decode converts the binary string back to original using the Huffman tree
func decode(encoded string, root *Node) string {
	var result strings.Builder
	current := root

	for _, bit := range encoded {
		if bit == '0' {
			current = current.left
		} else {
			current = current.right
		}

		if current.char != 0 {
			result.WriteRune(current.char)
			current = root
		}
	}

	return result.String()
}

func main() {
	data := "hello world"

	// Step 1: Build frequency map
	freq := buildFrequencyMap(data)

	// Step 2: Build Huffman tree
	root := buildHuffmanTree(freq)

	// Step 3: Generate code map
	codeMap := make(map[rune]string)
	buildCodeMap(root, "", codeMap)

	// Step 4: Encode the data
	encoded := encode(data, codeMap)

	// Step 5: Decode back to verify
	decoded := decode(encoded, root)

	fmt.Println("Original data:", data)
	fmt.Println("Frequency map:", freq)
	fmt.Println("Huffman codes:")
	for char, code := range codeMap {
		fmt.Printf("'%c': %s\n", char, code)
	}
	fmt.Println("Encoded binary:", encoded)
	fmt.Println("Decoded data:", decoded)
	fmt.Println("Compression ratio:", float64(len(data)*8)/float64(len(encoded)))
}
