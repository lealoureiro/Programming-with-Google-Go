package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
