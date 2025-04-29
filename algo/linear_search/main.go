package main

import (
	"fmt"
	"math/rand"
)

type Array struct {
	arr []int
}

func NewArray() *Array {
	return &Array{}
}

func (arr *Array) Add(val int) {
	arr.arr = append(arr.arr, val)
}

func (arr *Array) linearSearch(value int) int {
	for index, val := range arr.arr {
		if val == value {
			return index
		}
	}
	return -1
}

func main() {

	arr := NewArray()
	for i := 0; i < 100; i++ {
		arr.Add(rand.Intn(100))
	}
	index := arr.linearSearch(55)
	fmt.Println(index)
}
