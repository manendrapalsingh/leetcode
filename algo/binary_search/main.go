package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Array struct {
	arr []int
}

func NewArray() *Array {
	return &Array{}
}

func (arr *Array) RendomArray() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		arr.arr = append(arr.arr, rand.Intn(100))
	}
}

func (arr *Array) SortArray() {

	sort.Slice(arr.arr, func(i int, j int) bool {
		return arr.arr[i] < arr.arr[j] // to comapre in any condition
	})
}

func (arr *Array) BinarySearchRecusive(start int, end int, find int) int {

	if end < start {
		return -1
	}
	mid := (start + end) / 2

	if arr.arr[mid] == find {
		return mid
	} else if arr.arr[mid] < find {

		start = mid + 1

		return arr.BinarySearchRecusive(start, end, find)

	} else {
		end = mid - 1
		return arr.BinarySearchRecusive(start, end, find)
	}

}

func (arr *Array) BinarySearchIter(start int, end int, find int) int {

	for start <= end {
		mid := (start + end) / 2
		if arr.arr[mid] == find {
			return mid
		} else if arr.arr[mid] < find {

			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return -1
}
func main() {

	arr := NewArray()

	arr.RendomArray()

	arr.SortArray()

	fmt.Println(arr.arr)

	fmt.Println(arr.BinarySearchIter(0, len(arr.arr)-1, 15))
	fmt.Println(arr.BinarySearchRecusive(0, len(arr.arr)-1, 15))

}
