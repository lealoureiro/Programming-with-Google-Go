package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Printf("eats: %s\n", a.food)
}

func (a *Animal) Move() {
	fmt.Printf("moves: %s\n", a.locomotion)
}

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
