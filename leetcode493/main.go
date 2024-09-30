package main

import "fmt"

// Given an integer array nums, return the number of reverse pairs in the array.

// A reverse pair is a pair (i, j) where:

// 0 <= i < j < nums.length and
// nums[i] > 2 * nums[j].

// Example 1:

// Input: nums = [1,3,2,3,1]
// Output: 2
// Explanation: The reverse pairs are:
// (1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
// (3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1
// Example 2:

// Input: nums = [2,4,3,5,1]
// Output: 3
// Explanation: The reverse pairs are:
// (1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
// (2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
// (3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1

func main() {

	nums := []int{2, 3, 4, 5, 1}

	solution := 2 // please select the solution which you want to executed

	if solution == 1 {
		goto solution1
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

		count := mergeSort(nums, 0, len(nums)-1)

		fmt.Println(count)

		goto end
	}

end:
	fmt.Println("Execution Completed.")

}

func countPair(arr []int, low int, mid int, high int) int {
	right := mid + 1

	count := 0

	for i := low; i <= mid; i++ {

		for right <= high && arr[i] > 2*arr[right] {
			right += 1
		}
		count += right - (mid + 1)

	}

	return count

}

func mergeSort(arr []int, low int, high int) int {

	count := 0

	if low >= high {
		return count
	}

	mid := (low + high) / 2

	count += mergeSort(arr, low, mid)
	count += mergeSort(arr, mid+1, high)
	count += countPair(arr, low, mid, high)

	merge(arr, low, mid, high)

	return count
}

func merge(arr []int, low int, mid int, high int) []int {

	temArr := []int{}

	left := low
	right := mid + 1

	for left <= mid && right <= high {

		if arr[left] <= arr[right] {
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
