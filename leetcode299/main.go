package main

import "fmt"

// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

// Example 1:

// Input: nums = [3,2,3]
// Output: [3]
// Example 2:

// Input: nums = [1]
// Output: [1]
// Example 3:

// Input: nums = [1,2]
// Output: [1,2]

// Constraints:

// 1 <= nums.length <= 5 * 104
// -109 <= nums[i] <= 109

func main() {

	nums := []int{3}

	lenOfNums := len(nums)

	frequency := make(map[int]int)

	Output := []int{}

	for _, value := range nums {

		if _, ok := frequency[value]; ok {

			frequency[value] += 1

		} else {

			frequency[value] = 1
		}
	}

	for key, value := range frequency {

		if value > lenOfNums/3 {

			Output = append(Output, key)

		}
	}
	fmt.Println(Output)
}
