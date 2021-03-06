// Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

// s =½ a t2 + vot + so

// Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

// You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so.
// GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

// For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1.
// I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

// fn := GenDisplaceFn(10, 2, 1)

// Then I can use the following statement to print the displacement after 3 seconds.

// fmt.Println(fn(3))

// And I can use the following statement to print the displacement after 5 seconds.

// fmt.Println(fn(5))

package main

import (
	"fmt"
)

func main() {

	a, v0, s0 := getUserValues()

	fmt.Println("Input values:")
	fmt.Printf("acceleration: %.2f\n", a)
	fmt.Printf("initial velocity: %.2f\n", v0)
	fmt.Printf("initial displacement: %.2f\n", s0)

	fn := genDisplaceFn(a, v0, s0)

	fmt.Printf("displacement after 3 seconds: %.2f\n", fn(3))
	fmt.Printf("displacement after 5 seconds: %.2f\n", fn(5))

}

func getUserValues() (float64, float64, float64) {

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

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}

}
