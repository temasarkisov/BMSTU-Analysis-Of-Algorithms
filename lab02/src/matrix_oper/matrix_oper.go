package matrixoper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// OutputMatrix output matrix in apropriate style
func OutputMatrix(matrix [][]int) {
	for i := range matrix {
		for _, value := range matrix[i] {
			fmt.Print(value)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

// InputMatrix allows user to enter matrix
func InputMatrix(matrix [][]int) [][]int {
	intSlice := []int{} // Slice containing entered integers
	rowN := 0
	reader := bufio.NewReader(os.Stdin)

	for {
		inputString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if inputString == "\n" {
			return matrix
		}
		inputString = (inputString[:len(inputString)-1])
		if inputString[len(inputString)-1] == ' ' { // Preventing space as last symbol
			inputString = (inputString[:len(inputString)-1])
		}
		stringSlice := strings.Split(inputString, " ")
		for _, strValue := range stringSlice {
			intValue, err := strconv.Atoi(strValue)
			if err != nil {
				panic(err)
			}
			intSlice = append(intSlice, intValue)
		}
		matrix = append(matrix, intSlice)
		intSlice = intSlice[:0]
		rowN++
	}
	return matrix
}

// CreateMatrix creates matrix based on rows and columns number
func CreateMatrix(rowsNum int, colsNum int) [][]int {
	matrix := make([][]int, rowsNum)
	for i := range matrix {
		matrix[i] = make([]int, colsNum)
	}
	return matrix
}

// MatrixMultStd multiplies two matrices based on standart algorithm
func MatrixMultStd(matrix1 [][]int, matrix2 [][]int) [][]int {
	var n1 int = len(matrix1)
	var n2 int = len(matrix2)

	if n1 == 0 || n2 == 0 {
		return nil
	}

	var m1 int = len(matrix1[0])
	var m2 int = len(matrix2[0])

	if m1 != n2 {
		return nil
	}

	result := CreateMatrix(n1, m2)

	for i := 0; i < n1; i++ {
		for j := 0; j < m2; j++ {
			for k := 0; k < m1; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}

	return result
}

// MatrixMultVin multiplies two matrices based on Vinograd algorithm
func MatrixMultVin(matrix1 [][]int, matrix2 [][]int) [][]int {
	var n1 int = len(matrix1)
	var n2 int = len(matrix2)

	if n1 == 0 || n2 == 0 {
		return nil
	}

	var m1 int = len(matrix1[0])
	var m2 int = len(matrix2[0])

	if m1 != n2 {
		return nil
	}

	mulH := make([]int, n1)
	mulV := make([]int, m2)
	result := CreateMatrix(n1, m2)

	for i := 0; i < n1; i++ {
		for j := 0; j < m1/2; j++ {
			mulH[i] += matrix1[i][j*2] * matrix1[i][j*2+1]
		}
	}

	for i := 0; i < m2; i++ {
		for j := 0; j < n2/2; j++ {
			mulV[i] += matrix2[j*2][i] * matrix2[j*2+1][i]
		}
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < m2; j++ {
			result[i][j] = -mulH[i] - mulV[j]
			for k := 0; k < m1/2; k++ {
				result[i][j] += (matrix1[i][2*k+1] + matrix2[2*k][j]) * (matrix1[i][2*k] + matrix2[2*k+1][j])
			}
		}
	}

	if m1%2 == 1 {
		for i := 0; i < n1; i++ {
			for j := 0; j < m2; j++ {
				result[i][j] += matrix1[i][m1-1] * matrix2[m1-1][j]
			}
		}
	}

	return result
}

// MatrixMultVinOptim multiplies two matrices based on Vinograd algorithm with optimization
func MatrixMultVinOptim(matrix1 [][]int, matrix2 [][]int) [][]int {
	var n1 int = len(matrix1)
	var n2 int = len(matrix2)

	if n1 == 0 || n2 == 0 {
		return nil
	}

	var m1 int = len(matrix1[0])
	var m2 int = len(matrix2[0])

	if m1 != n2 {
		return nil
	}

	mulH := make([]int, n1)
	mulV := make([]int, m2)
	result := CreateMatrix(n1, m2)

	var m1Mod2 int = m1 % 2
	var n2Mod2 int = n2 % 2

	for i := 0; i < n1; i++ {
		for j := 0; j < m1-m1Mod2; j += 2 {
			mulH[i] += matrix1[i][j] * matrix1[i][j+1]
		}
	}

	for i := 0; i < m2; i++ {
		for j := 0; j < n2-n2Mod2; j += 2 {
			mulV[i] += matrix2[j][i] * matrix2[j+1][i]
		}
	}

	var buff int
	for i := 0; i < n1; i++ {
		for j := 0; j < m2; j++ {
			buff = -mulH[i] - mulV[j]
			for k := 0; k < m1-m1Mod2; k += 2 {
				buff += (matrix1[i][k+1] + matrix2[k][j]) * (matrix1[i][k] + matrix2[k+1][j])
			}
			result[i][j] = buff
		}
	}

	if m1Mod2 == 1 {
		var m1Min1 int = m1 - 1
		for i := 0; i < n1; i++ {
			for j := 0; j < m2; j++ {
				result[i][j] += matrix1[i][m1Min1] * matrix2[m1Min1][j]
			}
		}
	}

	return result
}
