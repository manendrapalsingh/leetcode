package main

import (
	"fmt"
	"slices"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.


Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

func merge(intervals [][]int) [][]int {

	slices.SortFunc(intervals, func(i, j []int) int {
		if i[0] == j[0] {
			return i[1] - j[1] // ascending by end
		}
		return i[0] - j[0] // ascending by start
	})

	fmt.Println(intervals)
	for i := 0; i < len(intervals)-1; i++ {
		j := i + 1

		if slices.Max(intervals[i]) >= slices.Min(intervals[j]) && slices.Min(intervals[i]) <= slices.Max(intervals[j]) {

			var tmpArr []int

			tmpArr = append(tmpArr, intervals[i]...)
			tmpArr = append(tmpArr, intervals[j]...)

			intervals[j] = []int{slices.Min(tmpArr), slices.Max(tmpArr)}

			intervals = slices.Delete(intervals, i, j)
			i++
			j++

		}

	}

	return intervals

}

func main() {

	input := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}

	merge(input)
}
