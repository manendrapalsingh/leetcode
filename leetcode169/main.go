package main

import "fmt"

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// Constraints:

// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109

func main() {

	nums := []int{3, 2, 3}

	freq := make(map[int]int)

	var maxfreq int
	var maxfreqInt int

	for _, val := range nums {

		if _, ok := freq[val]; ok {

			freq[val] += 1
		} else {
			freq[val] = 1
		}

	}

	for key, val := range freq {

		if val > maxfreq {

			maxfreq = val
			maxfreqInt = key

		}
	}

	fmt.Println(maxfreqInt, maxfreq)
}
