package main

import (
	"fmt"
)

func main() {

	a, v0, s0 := GetUserValues()

	fmt.Println("Input values:")
	fmt.Printf("acceleration: %.2f\n", a)
	fmt.Printf("initial velocity: %.2f\n", v0)
	fmt.Printf("initial displacement: %.2f\n", s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Printf("displacement after 3 seconds: %.2f\n", fn(3))
	fmt.Printf("displacement after 5 seconds: %.2f\n", fn(5))

}

func GetUserValues() (float64, float64, float64) {

	var a, v0, s0 float64

	fmt.Printf("Please insert acceleration: ")
	fmt.Scan(&a)

	fmt.Printf("Please insert initial velocity: ")
	fmt.Scan(&v0)

	fmt.Printf("Please insert initial displacement: ")
	fmt.Scan(&s0)

	fmt.Println()

	return a, v0, s0

}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}

}
