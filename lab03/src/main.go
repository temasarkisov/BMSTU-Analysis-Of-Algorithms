package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./sorts"
)

func inputSlice() []int {
	intSlice := []int{} // Slice containing entered integers
	fmt.Print("Enter series of integers: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
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
	return intSlice
}

func printSlice(intSlice []int) {
	for _, value := range intSlice {
		fmt.Printf("%d ", value)
	}
	fmt.Print("\n")
}

func main() {
	intSlice := inputSlice()
	//sorts.BubbleSort(intSlice)
	//sorts.InsertionSort(intSlice)
	sorts.QuickSort(intSlice)
	printSlice(intSlice)

}
