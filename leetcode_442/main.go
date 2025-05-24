package main

import "fmt"

/*

Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears at most twice, return an array of all the integers that appears twice.

You must write an algorithm that runs in O(n) time and uses only constant auxiliary space, excluding the space needed to store the output



Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [2,3]
Example 2:

Input: nums = [1,1,2]
Output: [1]
Example 3:

Input: nums = [1]
Output: []

*/

func main() {

	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	output := make([]int, 0)

	hash := make(map[int]bool)

	for _, v := range input {

		if _, ok := hash[v]; ok {
			output = append(output, v)
		} else {
			hash[v] = true
		}
	}

	fmt.Println(output)
}
