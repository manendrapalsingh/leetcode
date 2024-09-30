package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

func main() {

	nums := []int{2, 3, 4, 5, 1}

	solution := 2 // please select the solution which you want to executed

	if solution == 1 {
		goto solution1 // O(n^2)
	} else if solution == 2 {
		goto solution2
	}

	// to long to get execute

solution1:
	{

		// here we simply loop through the arr and do the another nested loop to check the condition
		numbersOfPair := 0
		lenOfNums := len(nums)

		for i := 0; i < lenOfNums; i++ {

			for j := i + 1; j < lenOfNums; j++ {

				if nums[i] > 2*nums[j] {

					numbersOfPair += 1
				}

			}
		}
		fmt.Println(numbersOfPair)
		goto end
	}

	// optimized version

solution2:
	{
		// here we are going to use merge short but in different ways

		desArr := mergeSort(nums, 0, len(nums)-1)

		fmt.Println(desArr)

		goto end
	}

end:
	fmt.Println("Execution Completed.")

}

func mergeSort(arr []int, low int, high int) []int {

	if low >= high {
		return nil
	}

	mid := (low + high) / 2

	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)

	return merge(arr, low, mid, high)
}

func merge(arr []int, low int, mid int, high int) []int {

	temArr := []int{}

	left := low
	right := mid + 1

	for left <= mid && right <= high {

		if arr[left] >= arr[right] {
			temArr = append(temArr, arr[left])
			left += 1
		} else {

			temArr = append(temArr, arr[right])
			right += 1
		}
	}

	for left <= mid {
		temArr = append(temArr, arr[left])
		left += 1
	}

	for right <= high {
		temArr = append(temArr, arr[right])
		right += 1

	}

	for i := low; i <= high; i++ {
		arr[i] = temArr[i-low]
	}

	return arr
}
