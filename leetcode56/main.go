package main

import (
	"fmt"
	"sort"
)

// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
// Example 2:

// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

func main() {

	inputArr := [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}

	sort.Slice(inputArr, func(i, j int) bool {

		if inputArr[i][0] != inputArr[j][0] {

			return inputArr[i][0] < inputArr[j][0]
		}

		return inputArr[i][1] < inputArr[j][1]

	})

	fmt.Println(inputArr)

	lenOfArr := len(inputArr)

	outputArr := [][]int{}

	var j int

	for i := 0; i < lenOfArr; i++ {

		fmt.Println(outputArr, j, i)

		if (i < lenOfArr-1) && (inputArr[i][1] >= inputArr[i+1][0]) && ((len(outputArr) > 0) && (outputArr[j-1][1] < inputArr[i][0])) {

			outputArr = append(outputArr, []int{min(inputArr[i][0], inputArr[i+1][0]), max(inputArr[i][1], inputArr[i+1][1])})
			j++
			i++
		} else if (i < lenOfArr) && (len(outputArr) > 0) && (outputArr[j-1][1] >= inputArr[i][0]) {

			outputArr[j-1][0] = min(outputArr[j-1][0], inputArr[i][0])
			outputArr[j-1][1] = max(outputArr[j-1][1], inputArr[i][1])

		} else {
			outputArr = append(outputArr, inputArr[i])
			j++
		}

	}

	fmt.Println(outputArr)
}
