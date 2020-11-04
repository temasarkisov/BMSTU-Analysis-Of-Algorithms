package main

import (
	"fmt"

	matrixoper "./matrix_oper"
)

func main() {
	var operNum int = -1
	var menuFlag bool = true
	matrix1 := [][]int{}
	matrix2 := [][]int{}
	result := [][]int{}

	for menuFlag != false {
		fmt.Print("1 - Ввод матриц\n")
		fmt.Print("2 - Перемножение матриц стандартным методом\n")
		fmt.Print("3 - Перемножение матриц методом Винограда\n")
		fmt.Print("4 - Перемножение матриц оптимизированным методом Винограда\n")
		fmt.Print("0 - Выход\n")

		fmt.Print("\n>")
		_, err := fmt.Scanf("%d", &operNum)
		if err != nil {
			fmt.Print("Неверный номер операции!\n")
			panic(err)
		}

		switch operNum {
		case 0:
			fmt.Print("До свидания!\n")
			menuFlag = false
		case 1:
			fmt.Print("Введите первую матрицу построчно:\n")
			matrix1 = matrixoper.InputMatrix(matrix1)
			fmt.Print("Введите вторую матрицу построчно:\n")
			matrix2 = matrixoper.InputMatrix(matrix2)
		case 2:
			result = matrixoper.MatrixMultStd(matrix1, matrix2)
			fmt.Print("Полученная матрица:\n")
			matrixoper.OutputMatrix(result)
		case 3:
			result = matrixoper.MatrixMultStd(matrix1, matrix2)
			fmt.Print("Полученная матрица:\n")
			matrixoper.OutputMatrix(result)
		case 4:
			result = matrixoper.MatrixMultVinOptim(matrix1, matrix2)
			fmt.Print("Полученная матрица:\n")
			matrixoper.OutputMatrix(result)
		default:
			fmt.Print("Неверный номер операции!\n")
		}
	}
}
