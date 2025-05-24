package main

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	data [][]int
}

func NewMatrix() *Matrix {
	return &Matrix{}
}

func (matrix *Matrix) InsertIntoMAtrix(m int, n int) {

	matrix.data = make([][]int, m)

	for i := 0; i < m; i++ {

		matrix.data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix.data[i][j] = rand.Intn(100)
		}
	}

}

// formaula for matrix multiplacation Cij = sum from k=0 to k=n  Aik * Bkj
func NormalMatrixMultipaction(matrixA [][]int, matrixB [][]int) ([][]int, error) {

	// wg := sync.WaitGroup{}

	colsA := len(matrixA[0])
	rowsB := len(matrixB)
	if colsA != rowsB {
		return nil, fmt.Errorf("matrix dimensions incompatible for multiplication")
	}

	rowsA := len(matrixA)
	colsB := len(matrixB[0])

	output := make([][]int, rowsA)
	for i := range output {
		output[i] = make([]int, colsB)
	}

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			for k := 0; k < colsA; k++ {
				// wg.Add(1)
				// go func(wg *sync.WaitGroup) {

				// defer wg.Done()
				output[i][j] += matrixA[i][k] * matrixB[k][j]

				// }(&wg)
			}
		}
	}

	// wg.Wait()
	return output, nil

}

func main() {

	matrix1 := NewMatrix()
	matrix2 := NewMatrix()

	matrix1.InsertIntoMAtrix(1000, 1000)
	matrix2.InsertIntoMAtrix(1000, 1000)

	_, _ = NormalMatrixMultipaction(matrix1.data, matrix2.data)
}
