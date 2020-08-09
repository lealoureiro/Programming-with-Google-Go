// Write a program which allows the user to get information about a predefined set of animals.
// Three animals are predefined, cow, bird, and snake.
// Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal:
//1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
//The following table contains the three animals and their associated data which should be hard-coded into your program.

//Animal	Food eaten	Locomotion method	Spoken sound
//cow		grass		walk				moo
//bird		worms		fly					peep
//snake		mice		slither				hsss

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
// Your program should continue in this loop forever.
// Every request from the user must be a single line containing 2 strings.
// The first string is the name of an animal, either “cow”, “bird”, or “snake”.
// The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
// Your program should process each request by printing out the requested data.

// You will need a data structure to hold the information about each animal.
// Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings.
// Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
// Your program should call the appropriate method when the user makes a request.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal holds the animal properties
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat prints what animal eats
func (a *Animal) Eat() {
	fmt.Printf("eats: %s\n", a.food)
}

// Move print how animal Moves
func (a *Animal) Move() {
	fmt.Printf("moves: %s\n", a.locomotion)
}

// Speak prints how animal makes noise
func (a *Animal) Speak() {
	fmt.Printf("makes noise: %s\n", a.noise)
}

var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fly", "peep"}
var snake = Animal{"mice", "slither", "hsss"}

func main() {

	fmt.Println("Animals")

	var terminated bool = false
	var printInformation func()

	for !terminated {

		printInformation, terminated = getUserRequest()

		printInformation()

	}

}

func getUserRequest() (func(), bool) {

	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	request := strings.TrimSpace(strings.ToLower(scanner.Text()))

	requestParts := strings.Split(request, " ")

	if len(requestParts) == 0 {
		return func() {
			fmt.Println("Insert your request with following format: <animal> <information> or \"exit\" to quit!")
		}, false
	}

	if len(requestParts) == 1 && requestParts[0] == "exit" {
		return func() { fmt.Println("Leaving application!") }, true
	}

	if len(requestParts) == 2 {
		return processRequest(requestParts[0], requestParts[1]), false
	}

	return func() {
		fmt.Println("Insert your request with following format: <animal> <information> or \"exit\" to quit!")
	}, false
}

func processRequest(animal, information string) func() {

	switch animal {
	case "cow":
		return processInformation(cow, animal, information)
	case "bird":
		return processInformation(bird, animal, information)
	case "snake":
		return processInformation(snake, animal, information)
	default:
		return func() { fmt.Println("Invalid animal! Availaible options: cow, bird, snake.") }
	}

}

func processInformation(animal Animal, animalName, information string) func() {
	switch information {
	case "eat":
		return func() {
			fmt.Printf("Animal %s ", animalName)
			animal.Eat()
		}
	case "move":
		return func() {
			fmt.Printf("Animal %s ", animalName)
			animal.Move()
		}
	case "speak":
		return func() {
			fmt.Printf("Animal %s ", animalName)
			animal.Speak()
		}
	default:
		return func() {
			fmt.Printf("Requested invalid information for animal %s! available options: eat, move, speak!\n", animalName)
		}
	}
}
