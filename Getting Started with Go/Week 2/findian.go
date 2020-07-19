package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Please type a string:")

	scanner := bufio.NewScanner(os.Stdin)

	var normalizedString string = ""

	scanner.Scan()
	normalizedString = strings.ToLower(scanner.Text())

	if strings.HasPrefix(normalizedString, "i") && strings.Contains(normalizedString, "a") && strings.HasSuffix(normalizedString, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}

}
