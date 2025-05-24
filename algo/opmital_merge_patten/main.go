package main

import (
	"sort"
)

func main() {

	// list := []string{"x1", "x2", "x3", "x4", "x5"}
	size := []int{20, 30, 10, 5, 30}

	sort.Slice(size, func(i, j int) bool {
		return size[i] < size[j]
	})

	// var newMergedList []int
	totalCost := 0

	for len(size) != 1 {

		if len(size) >= 2 {
			tmpSum := size[0] + size[1]
			size[0] = tmpSum
			size = append(size, size[2:]...)
			totalCost += tmpSum

			sort.Slice(size, func(i, j int) bool {
				return size[i] < size[j]
			})

		} else {
			break
		}

	}
}
