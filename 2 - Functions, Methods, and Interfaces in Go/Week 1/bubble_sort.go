//Write a Bubble Sort program in Go.
// The program should prompt the user to type in a sequence of up to 10 integers.
// The program should print the integers out on one line, in sorted order, from least to greatest.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.

// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.

// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
// You should write a Swap() function which performs this operation.
// Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
// The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Bubble Sort")

	numbers := getNumbersFromUser()

	printNumbers(numbers)

	bubbleSort(numbers)

	printSortedNumbers(numbers)

}

func getNumbersFromUser() []int {

	unsortedNumbers := make([]int, 0, 10)
	var data string
	finished := false
	count := 0

	for !finished && count < 10 {

		fmt.Printf("Insert number %d: ", count+1)
		fmt.Scan(&data)

		if strings.Compare(data, "C") == 0 {

			finished = true

		} else {

			number, err := strconv.Atoi(data)

			if err == nil {

				unsortedNumbers = append(unsortedNumbers, number)
				count++

			} else {

				fmt.Println("Insert a correct number or 'C' to complete!")

			}

		}

	}

	return unsortedNumbers
}

func printNumbers(numbers []int) {

	fmt.Print("You insert the following numbers: ")

	for _, n := range numbers {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}

func printSortedNumbers(numbers []int) {

	fmt.Print("The list of numbers sorted is: ")

	for _, n := range numbers {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}

func bubbleSort(numbers []int) {

	totalNumbers := len(numbers)

	for i := 0; i < totalNumbers-1; i++ {

		for j := 0; j < totalNumbers-i-1; j++ {

			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}

		}

	}

}

func swap(numbers []int, index int) {

	number := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = number

}
