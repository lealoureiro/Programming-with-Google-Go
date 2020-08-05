package main

import "fmt"

/*
	Because we set the value of i in a different Go rotine we don't know if the return value that will be printed will happen before
	or after the instruction to set the value to 3, this can be a race condition
*/
func main() {
	fmt.Println(getNumberTree())
}

func getNumberTree() int {
	var i int

	go setNumberTree(&i)

	return i
}

func setNumberTree(n *int) {
	*n = 3
}
