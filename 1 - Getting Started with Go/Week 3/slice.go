package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var data string
	var s1 = make([]int, 0, 3)
	finished := false

	for !finished {

		fmt.Printf("Insert a number: ")
		fmt.Scan(&data)

		if strings.Compare(data, "X") == 0 {

			finished = true

		} else {

			number, err := strconv.Atoi(data)

			if err == nil {

				s1 = append(s1, number)
				sort.Ints(s1)

			} else {

				fmt.Println("Insert a correct number or 'X' to exit!")

			}

		}

		fmt.Printf("Numbers: %d\n", s1)

	}

}
