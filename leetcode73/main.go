package main

import "fmt"

// Question Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// approach
//1 -> create a place holder 2d array to save the coordinates of zero
//2 -> loop through this place holder and make all the row and col corresponding to the coordinate to zero

func main() {

	inputMatrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}

	zeroAt := [][]int{}

	for i, v := range inputMatrix {

		for j, val := range v {

			if val == 0 {

				fmt.Println(i, j)

				zeroAt = append(zeroAt, []int{i, j})

			}

		}

	}

	for _, v := range zeroAt {

		lenOfRow := len(inputMatrix)
		lenOfColumn := len(inputMatrix[0])

		for r := 0; r < lenOfRow; r++ {

			inputMatrix[r][v[1]] = 0

		}
		for c := 0; c < lenOfColumn; c++ {

			inputMatrix[v[0]][c] = 0
		}

	}

	fmt.Println(inputMatrix)

}
