package main

import "fmt"

// Input:      {{4,5,6},
//              {7,8,9},
//              {10,11,12}}
// Output:
//               4 7 10
//               5 8 11
//               6 9 12

func main() {

	inputMatrix := [][]int{{4, 5, 6}, {7, 8, 9}, {10, 11, 12}}

	// for the transpose of the input matrix we can assume the diagonal(i=j) as mirror and take a reflection matrix

	for i := 0; i < len(inputMatrix); i++ {

		for j := 0; j < i; j++ {

			inputMatrix[i][j], inputMatrix[j][i] = inputMatrix[j][i], inputMatrix[i][j]

		}
	}

	fmt.Println(inputMatrix)

}
