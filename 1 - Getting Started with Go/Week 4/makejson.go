// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var m map[string]string = make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Insert your name: ")
	scanner.Scan()
	m["name"] = scanner.Text()

	fmt.Printf("Insert your address: ")
	scanner.Scan()
	m["address"] = scanner.Text()

	data, _ := json.Marshal(m)

	fmt.Printf("JSON: %s\n", data)

}
