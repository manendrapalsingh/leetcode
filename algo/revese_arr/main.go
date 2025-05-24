package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Array struct {
	data []int
}

func NewArray(size int) *Array {
	return &Array{data: make([]int, size)}
}

func (a *Array) Create() {

	for i := 0; i < len(a.data); i++ {
		a.data[i] = rand.Intn(100)
	}
}

func Merge(array []int, array2 []int) []int {

	output := make([]int, len(array)+len(array2))
	i, j, k := 0, 0, 0

	for i < len(array) && j < len(array2) {
		if array[i] < array2[j] {
			output[k] = array[i]
			i++

		} else {
			output[k] = array2[j]
			j++
		}
		k++
	}

	for i < len(array) {
		output[k] = array[i]
		k++
		i++
	}

	for j < len(array2) {
		output[k] = array2[j]
		k++
		j++
	}

	return output

}
func (a *Array) Reverse() {

	i := 0
	j := len(a.data) - 1

	for i < j {
		a.data[i], a.data[j] = a.data[j], a.data[i]
		i++
		j--
	}

}

func Union(array, array2 []int) []int {

	// if you want to size of output array to be dynamic then use append operation
	output := make([]int, len(array)+len(array2))

	i, j, k := 0, 0, 0

	for i < len(array) && j < len(array2) {

		if array[i] == array[j] {
			output[k] = array[i]
			i++
			j++

		} else if array[i] < array2[j] {
			output[k] = array[i]
			i++

		} else {
			output[k] = array2[j]
			j++
		}
		k++
	}

	for i < len(array) {
		output[k] = array[i]
		k++
		i++
	}

	for j < len(array2) {
		output[k] = array2[j]
		k++
		j++
	}

	fmt.Println(len(output))
	return output
}

func Difference(array, array2 []int) []int {

	output := make([]int, 0)
	i, j := 0, 0

	for i < len(array) && j < len(array2) {

		if array[i] == array[j] {
			i++
			j++

		} else if array[i] < array2[j] {
			output = append(output, array[i])
			i++

		} else {
			j++
		}
	}

	output = append(output, array[i:]...)

	return output

}
func main() {

	array := NewArray(10)
	array.Create()

	//fmt.Println(array.data)
	array.Reverse()
	//fmt.Println(array.data)

	sort.Slice((array.data), func(i, j int) bool {
		return array.data[i] < array.data[j]
	})

	fmt.Println(array.data)

	array2 := NewArray(10)
	array2.Create()
	sort.Slice(array2.data, func(i, j int) bool {
		return array2.data[i] < array2.data[j]
	})

	fmt.Println(array2.data)

	fmt.Println(Merge(array.data, array2.data))

	fmt.Println(Union(array.data, array2.data))

	fmt.Println(Difference(array.data, array2.data))

}
