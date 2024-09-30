package main

import "fmt"

// You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

func main() {

	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	target := 3

	var hasTarget = false
	for _, arr := range matrix {

		if hasTarget {
			break
		}
		for _, value := range arr {

			if target < value {
				break
			}

			if value == target {
				hasTarget = true
				break
			}
		}
	}
	fmt.Println(hasTarget, "================================")
}
