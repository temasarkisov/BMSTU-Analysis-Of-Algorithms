package multest

import (
	"testing"

	matrixoper "../matrix_oper"
)

func BenchmarkMatrixMultStd(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		matrixoper.MatrixMultStd(matrix1, matrix2)
	}
}

func BenchmarkMatrixMultVin(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		matrixoper.MatrixMultVin(matrix1, matrix2)
	}
}

func BenchmarkMatrixMultVinOptim(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		matrixoper.MatrixMultVinOptim(matrix1, matrix2)
	}
}
