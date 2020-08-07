package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	unsortedNumbers := getNumbersFromUser()

	fmt.Print("You insert the following numbers: ")
	printNumbers(unsortedNumbers)

	chunks := getChuncks(unsortedNumbers)

	c := make(chan []int)

	go sortSlice(1, chunks[0], c)
	go sortSlice(2, chunks[1], c)
	go sortSlice(3, chunks[2], c)
	go sortSlice(4, chunks[3], c)

	s1 := <-c
	s2 := <-c
	s3 := <-c
	s4 := <-c

	var sorted []int

	sorted = append(sorted, s1...)
	sorted = append(sorted, s2...)
	sorted = append(sorted, s3...)
	sorted = append(sorted, s4...)

	sort.Ints(sorted)

	fmt.Print("Numbers sorted: ")
	printNumbers(sorted)

}

func sortSlice(n int, s []int, c chan []int) {

	fmt.Printf("Routine %d will sort: ", n)
	printNumbers(s)

	sort.Ints(s)

	fmt.Printf("Routine %d sorted: ", n)
	printNumbers(s)

	c <- s

}

func getChuncks(numbers []int) [][]int {

	var chuncks [][]int = make([][]int, 4)

	for i, n := range numbers {
		bucket := i % 4
		chuncks[bucket] = append(chuncks[bucket], n)
	}

	return chuncks

}

func getNumbersFromUser() []int {

	unsortedNumbers := make([]int, 0)
	var data string
	finished := false
	count := 0

	for !finished {

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

	for _, n := range numbers {
		fmt.Printf("%d ", n)
	}

	fmt.Println()
}
