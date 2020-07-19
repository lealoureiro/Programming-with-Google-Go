package main

import "fmt"

func main() {

	fmt.Println("Insert a floating point number:")

	var n float64

	fmt.Scan(&n)

	fmt.Printf("The number was truncated to: %d\n", int64(n))

}
