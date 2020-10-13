package main

import "testing"

func BenchmarkMatrixMultStd(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		matrixMultStd(matrix1, matrix2)
	}
}

func BenchmarkMatrixMultVin(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		matrixMultVin(matrix1, matrix2)
	}
}

func BenchmarkMatrixMultVinOptim(b *testing.B) {
	matrix1 := [][]int{{0, 1, 2}, {3, 4, 5}}
	matrix2 := [][]int{{0, 1}, {3, 4}, {6, 7}}
	for n := 0; n < b.N; n++ {
		matrixMultVinOptim(matrix1, matrix2)
	}
}
