package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Bubble Sort")

	numbers := GetNumbersFromUser()

	PrintNumbers(numbers)

	BubbleSort(numbers)

	PrintSortedNumbers(numbers)

}

func GetNumbersFromUser() []int {

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

func PrintNumbers(numbers []int) {

	fmt.Print("You insert the following numbers: ")

	for _, n := range numbers {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}

func PrintSortedNumbers(numbers []int) {

	fmt.Print("The list of numbers sorted is: ")

	for _, n := range numbers {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}

func BubbleSort(numbers []int) {

	totalNumbers := len(numbers)

	for i := 0; i < totalNumbers-1; i++ {

		for j := 0; j < totalNumbers-i-1; j++ {

			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}

		}

	}

}

func Swap(numbers []int, index int) {

	number := numbers[index]
	numbers[index] = numbers[index+1]
	numbers[index+1] = number

}
