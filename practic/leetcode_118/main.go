package main

import "fmt"

/*
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]

Constraints:

1 <= numRows <= 30
*/

func generate(numRows int) [][]int {

	output := [][]int{}

	for i := 0; i < numRows; i++ {

		temArr := make([]int, i+1)

		if i == 0 {

			output = append(output, []int{1})
		}

		if i == 1 {
			output = append(output, []int{1, 1})
			output[i] = []int{1, 1}
		}

		if i > 1 {

			for j, _ := range temArr {

				if j == 0 || j == i {
					temArr[j] = 1
				} else {

					temArr[j] = output[i-1][j] + output[i-1][j-1]

				}

			}

			output = append(output, temArr)
		}

	}

	return output

}

func main() {

	fmt.Println(generate(5))

}
