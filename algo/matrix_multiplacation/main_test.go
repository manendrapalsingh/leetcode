package main

import "testing"

func Banckmark(b *testing.B) {

	matrix1 := NewMatrix()
	matrix2 := NewMatrix()

	matrix1.InsertIntoMAtrix(5000, 5000)
	matrix2.InsertIntoMAtrix(5000, 5000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = NormalMatrixMultipaction(matrix1.data, matrix2.data)
	}

}
