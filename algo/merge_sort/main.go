package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Array struct {
	array []int
}

func NewArray() *Array {
	return &Array{}
}

func (arr *Array) Inster() {

	for i := 0; i < 10; i++ {
		arr.array = append(arr.array, rand.Intn(100))
	}
}

func (arr *Array) SortArray() {

	sort.Slice(arr.array, func(i, j int) bool {
		return arr.array[i] < arr.array[j]

	})

}

func MergedSort(list1 []int, list2 []int, m int, n int) []int {

	var mergedList []int

	i, j := 0, 0

	for i < m && j < n {

		if list1[i] < list2[j] {
			mergedList = append(mergedList, list1[i])
			i++
		} else {
			mergedList = append(mergedList, list2[j])
			j++
		}
	}

	mergedList = append(mergedList, list1[i:]...)

	mergedList = append(mergedList, list2[j:]...)

	return mergedList

}

func TowWayMergedSort(array []int) []int {

	if len(array) == 1 {
		return array
	}

	list := make([][]int, len(array))

	for i := range array {
		list[i] = []int{array[i]}
	}

	for len(list) > 1 {

		var mergedList [][]int

		for i := 0; i < len(list); i += 2 {

			if i+1 < len(list) {
				mergedList = append(mergedList, MergedSort(list[i], list[i+1], len(list[i]), len(list[i+1])))
			} else {
				mergedList = append(mergedList, list[i])
			}

		}

		list = mergedList

	}

	return list[0]

}

func (arr *Array) Merge(low int, mid int, high int) {

	left := make([]int, mid-low+1)
	right := make([]int, high-mid)

	copy(left, arr.array[low:mid+1])
	copy(right, arr.array[mid+1:high+1])

	i, j, k := 0, 0, low

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr.array[k] = left[i]
			i++
		} else {
			arr.array[k] = right[j]
			j++
		}
		k++

	}

	for i < len(left) {
		arr.array[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr.array[k] = right[j]
		j++
		k++
	}

}

func (arr *Array) MergedSortAccutal(low int, high int) {

	if low < high {

		mid := (low + high) / 2
		arr.MergedSortAccutal(low, mid)
		arr.MergedSortAccutal(mid+1, high)
		arr.Merge(low, mid, high)
	}
}
func main() {

	list1 := NewArray()
	list2 := NewArray()

	list1.Inster()
	list2.Inster()

	list1.SortArray()
	list2.SortArray()

	finalList := MergedSort(list1.array, list2.array, len(list1.array), len(list2.array))

	fmt.Println(finalList)

	twoWaySort := TowWayMergedSort(list1.array)
	fmt.Println(list1.array)
	fmt.Println(twoWaySort)

	fmt.Println(list1.array)
	list1.MergedSortAccutal(0, len(list1.array)-1)
	fmt.Println(list1.array)

}
