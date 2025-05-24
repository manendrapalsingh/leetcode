package main

import (
	"fmt"
	"slices"
)

/*
A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = [1,2,3], the following are all the permutations of arr: [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, then the next permutation of that array is the permutation that follows it in the sorted container. If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.



Example 1:

Input: nums = [1,2,3]
Output: [1,3,2]
Example 2:

Input: nums = [3,2,1]
Output: [1,2,3]
Example 3:

Input: nums = [1,1,5]
Output: [1,5,1]


Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 100
*/

func nextPermutation(nums []int) {

	maxVal := -1
	var index int

	lenOfNums := len(nums)

	if slices.IsSortedFunc(nums, func(x, y int) int {

		if x > y {
			return -1
		} else {
			return 1
		}
	}) {

		slices.Sort(nums)

		fmt.Println(nums)
		return
	}

	// first chekh from last index and stop where array stop in decanding order
	for i := lenOfNums - 1; i >= 0; i-- {

		if maxVal <= nums[i] {

			maxVal = nums[i]

		} else {
			index = i
			break
		}
	}

	for y := 1; y <= 100; y++ {

		flage := false
		val := nums[index] + y
		for i := index + 1; i < lenOfNums; i++ {

			if nums[i] == val {
				nums[i], nums[index] = nums[index], nums[i]
				flage = true
				break
			}
		}

		if flage {
			break
		}
	}

	slices.SortFunc(nums[index+1:], func(x, y int) int {

		if x > y {
			return 1
		} else {
			return -1
		}
	})

}
func main() {

	// input := []int{9, 3, 5, 4, 1}

	// nextPermutation(input)

	// input2 := []int{1, 1, 5}

	// nextPermutation(input2)

	// input3 := []int{9, 3, 5, 1, 1}

	// nextPermutation(input3)

	input4 := []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	nextPermutation(input4)
}
