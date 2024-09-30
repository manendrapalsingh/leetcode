package main

import "fmt"

//  short the given array by using merge sort

func main() {

	arr := []int{2, 3, 4, 5, 1}

	lenOfArr := len(arr)

	sortedArr := mergeSort(arr, 0, lenOfArr-1)

	fmt.Println(sortedArr)

}

func mergeSort(arr []int, low int, high int) []int {

	if low >= high {
		return nil
	}

	mid := (low + high) / 2

	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)

	return merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) []int {

	temArr := []int{}

	left := low
	right := mid + 1

	for left <= mid && right <= high {

		if arr[left] <= arr[right] {
			temArr = append(temArr, arr[left])
			left += 1
		} else {

			temArr = append(temArr, arr[right])
			right += 1
		}
	}

	for left <= mid {
		temArr = append(temArr, arr[left])
		left += 1
	}

	for right <= high {
		temArr = append(temArr, arr[right])
		right += 1

	}

	for i := low; i <= high; i++ {
		arr[i] = temArr[i-low]
	}

	return arr
}
