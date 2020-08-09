// Write a program which reads information from a file and represents it in a slice of structs.
// Assume that there is a text file which contains a series of names.
// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
// Each field will be a string of size 20 (characters).

// Your program should prompt the user for the name of the text file.
// Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
// Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
// After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Person holds the first and last name of a person.
type Person struct {
	fname string
	lname string
}

func main() {

	fmt.Print("Insert file name to read the names: ")
	inputScanner := bufio.NewScanner(os.Stdin)

	inputScanner.Scan()
	filename := strings.TrimSpace(inputScanner.Text())

	fmt.Printf("Reading file %s ...\n", filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", filename)
		os.Exit(1)
	}

	var persons []Person = make([]Person, 0, 10)
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {

		names := strings.Split(fileScanner.Text(), " ")

		p := Person{fname: names[0], lname: names[1]}

		persons = append(persons, p)

	}

	for i, p := range persons {
		fmt.Printf("%d: %s\n", i, p)
	}

	f.Close()

}
