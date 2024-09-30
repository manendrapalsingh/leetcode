package main

import (
	"fmt"
	"slices"
)

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

// Example 1:

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
// Example 2:

// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].
// Example 3:

// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

func main() {

	nums1 := []int{0}

	nums2 := []int{}

	lenOfNums2 := len(nums2)
	lenOfNums1 := len(nums1)

	var j int

	// to get the index of last non-zero value
	for i := 0; i < lenOfNums1; i++ {

		if nums1[i] != 0 {

			j = i
		}
	}

	fmt.Println(j)

	//  if all the value are zero then empty the whole array
	if j > 0 {

		nums1 = slices.Delete(nums1, j+1, lenOfNums1)
	} else {

		// otherwise delete section which is empty
		nums1 = slices.DeleteFunc(nums1, func(i int) bool {

			return i == 0
		})
	}

	fmt.Println(nums1)

	// merge two arrays

	for i := 0; i < lenOfNums2; i++ {

		nums1 = append(nums1, nums2[i])

	}

	// simply sort the merged arrays

	slices.Sort(nums1)

	fmt.Println(nums1)

}
