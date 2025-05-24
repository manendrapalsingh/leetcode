package main

import (
	"fmt"
	"math/rand"
)

type Array struct {
	Items []int
}

func NewArray() *Array {
	return &Array{}
}

func (arr *Array) InsertIntoArray() {

	for i := 0; i < 10; i++ {
		arr.Items = append(arr.Items, rand.Intn(100))
	}
}

func QucikSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	pivot := arr[0]

	var left, right []int

	for _, val := range arr[1:] {

		if val < pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}

	}

	left = QucikSort(left)
	right = QucikSort(right)

	return append(append(left, pivot), right...)

}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIdx := partition(arr, low, high)

		// Recursively sort the left and right partitions
		quickSortHelper(arr, low, pivotIdx-1)
		quickSortHelper(arr, pivotIdx+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose the last element as pivot
	i := low           // Index of smaller element

	for j := low; j < high; j++ {
		// If current element is <= pivot, swap it with arr[i]
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Place the pivot in its correct position
	arr[i], arr[high] = arr[high], arr[i]
	return i // Return the pivot index
}
func main() {

	arr1 := NewArray()
	arr1.InsertIntoArray()
	fmt.Println(arr1.Items)
}
