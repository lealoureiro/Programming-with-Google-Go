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
