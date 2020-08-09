// Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.

package main

import "fmt"

func main() {

	fmt.Println("Insert a floating point number:")

	var n float64

	fmt.Scan(&n)

	fmt.Printf("The number was truncated to: %d\n", int64(n))

}
