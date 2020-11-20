package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"./sorts"
)

func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

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
	var num int
	fmt.Print("\n")
	fmt.Print("Enter a number of array's elements: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}
	intSlice := generateRandomSlice(num)
	fmt.Print("The random slice of integers: ")
	printSlice(intSlice)
	fmt.Print("\n")

	//sorts.BubbleSort(intSlice)
	//sorts.InsertionSort(intSlice)
	sorts.QuickSort(intSlice)

	fmt.Print("Sorted array: ")
	printSlice(intSlice)
	fmt.Print("\n")
}
