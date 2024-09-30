package main

import (
	"fmt"
	"time"
)

// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

// Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

// The test cases are generated so that the answer will be less than or equal to 2 * 109.

// Example 1:

// Input: m = 3, n = 7
// Output: 28
// Example 2:

// Input: m = 3, n = 2
// Output: 3
// Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down

// Constraints:

// 1 <= m, n <= 100

func main() {
	m := 4
	n := 4

	startTime := time.Now().Second()
	// To calculate the number of unique paths
	total := uniquePath(m, n)

	endTime := time.Now().Second() - startTime
	fmt.Println("Total Unique Paths:", total)
	fmt.Println("Execution Time (in seconds):", endTime)
}

func uniquePath(m int, n int) int {
	// Create a memoization table to store results of subproblems
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m)
	}

	fmt.Println(memo, "memoization")
	// Start from the top-left corner (0, 0)
	return countPath(0, 0, m, n, memo)
}

func countPath(currentXIndex int, currentYIndex int, m int, n int, memo [][]int) int {

	fmt.Println(memo, "memo in recurtion")
	time.Sleep(time.Second * 1)
	// Base case: If the current coordinate is the bottom-right corner, return 1 (one valid path)
	if currentXIndex == n-1 && currentYIndex == m-1 {
		return 1
	}

	// Check if the current coordinate is out of bounds
	if currentXIndex >= n || currentYIndex >= m {
		return 0
	}

	// If the result has already been computed, return it from the memo table
	if memo[currentXIndex][currentYIndex] != 0 {
		return memo[currentXIndex][currentYIndex]
	}

	// Otherwise, calculate the number of paths by moving down and to the right
	memo[currentXIndex][currentYIndex] = countPath(currentXIndex+1, currentYIndex, m, n, memo) +
		countPath(currentXIndex, currentYIndex+1, m, n, memo)

	return memo[currentXIndex][currentYIndex]
}
