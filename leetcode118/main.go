package main

import (
	"fmt"
)

// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

//Approach represents
//1 -> create a first row if the output 2d array had length = 0
//2 -> make a empty [] with length = col number and assign 1 and last value to 1
//3-> loop through the new row and check where values are 0 and assigned those indices with the value for previousRow[sameIndex  ] + previousRow[sameIndex -1 ]
//4 -> append the new row to the output

func main() {

	numberOfRows := 5

	outputMatrix := [][]int{}

	for i := 0; i < numberOfRows; i++ {

		if len(outputMatrix) == 0 {
			outputMatrix = append(outputMatrix, []int{1})
		} else {

			newRow := make([]int, i+1)
			for j, _ := range outputMatrix[i-1] {

				if j == 0 {

					fmt.Println(i, j)

					newRow[j] = 1
					newRow[len(newRow)-1] = 1
				} else {

					for index, val := range newRow {

						if val == 0 {

							newRow[index] = outputMatrix[i-1][index] + outputMatrix[i-1][index-1]

						}

					}

				}

			}
			outputMatrix = append(outputMatrix, newRow)
		}
	}

	fmt.Println(outputMatrix)
}
