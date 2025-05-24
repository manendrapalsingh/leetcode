package main

import "fmt"

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

// Example 1:

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
// Example 2:

// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

// Constraints:

// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000
	
func main() {

	// the original matrix is in the order of
	/*    |----------------------> Row
	      | (0,0) (0,1) (0,2)
	      | (1,0) (1,1) (1,2)
	      | (2,0) (2,1) (2,2)
		  |
		  column




		  90 degrees rotation then the oder will be

		 column <-----------------------------------|
			                   (2,0) (1,0) (0,0)	|
							   (2,1) (1,2) (0,2) 	|
			                   (2,2) (1,2) (0,2)	|
													|
													row


		row are interchanged into columns and columns are interchanged row
		and indexing will be started from top right corner instead of top left corner

	*/

	// inputMatrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	inputMatrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}

	{
		numbersOfRows, numbersOfColumns := len(inputMatrix), len(inputMatrix)

		var outputMatrix [][]int

		for i := 0; i < numbersOfRows; i++ {

			temArr := []int{}

			for j := 0; j < numbersOfColumns; j++ {

				temArr = append(temArr, inputMatrix[numbersOfColumns-j-1][i])

			}

			outputMatrix = append(outputMatrix, temArr)
		}

		fmt.Println(outputMatrix, "solution 1")
	}

	// soultion 2  take a transpose of the input matrix and reverse all of its rows
	{

		// transpose of input matrix
		for i := 0; i < len(inputMatrix); i++ {

			for j := 0; j < i; j++ {

				inputMatrix[i][j], inputMatrix[j][i] = inputMatrix[j][i], inputMatrix[i][j]

			}
		}

		// need to reverse the rows of input matrix

		for _, v := range inputMatrix {

			fmt.Println(v, "before reverse")

			for j, _ := range v {

				if len(v)%2 == 0 && j < (len(v)/2) {

					v[j], v[len(v)-j-1] = v[len(v)-j-1], v[j]

				} else {

					if j < (len(v) / 2) {
						v[j], v[len(v)-j-1] = v[len(v)-j-1], v[j]

					}

				}

			}

			fmt.Println(v, "after reverse")

		}
	}

	fmt.Println(inputMatrix, "solution 2")

}
