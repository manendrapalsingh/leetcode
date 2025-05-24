package main

import (
	"fmt"
	"sort"
)

/*
iven an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
Example 2:

Input: nums = [1,1]
Output: [2]

Constraints:

n == nums.length
1 <= n <= 105
1 <= nums[i] <= n
*/

func BinarySearch(array []int, start int, end int, find int) bool {

	if end < start {
		return false
	}
	mid := (start + end) / 2

	if array[mid] == find {
		return true
	} else if array[mid] > find {
		return BinarySearch(array, start, mid-1, find)
	} else {
		return BinarySearch(array, mid+1, end, find)
	}

}
func main() {
	nums := []int{10, 2, 5, 10, 9, 1, 1, 4, 3, 7}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	length := len(nums) - 1

	output := make([]int, 0)

	for i, _ := range nums {

		if !BinarySearch(nums, 0, length, i+1) {

			output = append(output, i+1)
		}

	}

	fmt.Println(output)

}

func SecondWay(nums []int) []int {

	output := make([]int, 0)

	for _, val := range nums {
		index := abs(val) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	for i, val := range nums {

		if val > 0 {
			output = append(output, i+1)
		}
	}
	return output
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
