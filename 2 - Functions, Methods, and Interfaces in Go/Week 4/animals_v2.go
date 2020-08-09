// Write a program which allows the user to create a set of animals and to get information about those animals.
// Each animal has a name and can be either a cow, bird, or snake.
// With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created.
// Each animal has a unique name, defined by the user.
// Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake.
// The following table contains the three types of animals and their associated data.

// Animal	Food eaten	Locomtion method	Spoken sound
// cow		grass		walk				moo
// bird		worms		fly					peep
// snake	mice		slither				hsss

// Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
// Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line.
// Your program should continue in this loop forever.
// Every command from the user must be either a “newanimal” command or a “query” command.

// Each “newanimal” command must be a single line containing three strings.
// The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal.
// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

// Each “query” command must be a single line containing 3 strings.
// The first string is “query”. The second string is the name of the animal.
// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
// Your program should process each query command by printing out the requested data.

// Define an interface type called Animal which describes the methods of an animal.
// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
// Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
// When the user creates an animal, create an object of the appropriate type.
// Your program should call the appropriate method when the user issues a query command.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal holds the animal properties
type Animal interface {
	Eat()
	Move()
	Speak()
	Type() string
}

// Cow animal
type Cow struct{ eat, move, speak string }

// Bird animal
type Bird struct{ eat, move, speak string }

// Snake animal
type Snake struct{ eat, move, speak string }

// Eat prints how Cow eats
func (c Cow) Eat() {
	fmt.Printf("Cow eats %s\n", c.eat)
}

// Move print how Cow moves
func (c Cow) Move() {
	fmt.Printf("Cow moves %s\n", c.move)
}

// Speak prints how Cow makes noise
func (c Cow) Speak() {
	fmt.Printf("Cow makes noise %s\n", c.speak)
}

// Type return the animal type Cow
func (c Cow) Type() string {
	return "Cow"
}

// Eat prints how Bird eats
func (b Bird) Eat() {
	fmt.Printf("Bird eats %s\n", b.eat)
}

// Move print how Bird moves
func (b Bird) Move() {
	fmt.Printf("Bird moves %s\n", b.move)
}

// Speak prints how Bird makes noise
func (b Bird) Speak() {
	fmt.Printf("Bird makes noise %s\n", b.speak)
}

// Type return the animal type Bird
func (b Bird) Type() string {
	return "Bird"
}

// Eat prints how Snake eats
func (s Snake) Eat() {
	fmt.Printf("Snake eats %s\n", s.eat)
}

// Move print how Snake moves
func (s Snake) Move() {
	fmt.Printf("Snake moves %s\n", s.move)
}

// Speak prints how Snake makes noise
func (s Snake) Speak() {
	fmt.Printf("Snake makes noise %s\n", s.speak)
}

// Type return the animal type Snake
func (s Snake) Type() string {
	return "Snake"
}

var animals map[string]Animal = make(map[string]Animal)

func main() {

	fmt.Println("Animals")

	var terminated bool = false

	fmt.Println("To insert a new Animal type: \"newanimal\" <name> <\"cow\", \"bird\" or \"snake\">")
	fmt.Println("To query an Animal type: \"query\" <name> <\"eat\", \"move\" or \"speak\">")
	fmt.Println("Type \"exit\" to quit.")

	for !terminated {

		terminated = getUserRequest()
		printCurrentAnimals()

	}

}

func getUserRequest() bool {

	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	request := strings.TrimSpace(scanner.Text())

	requestParts := strings.Split(request, " ")

	if len(requestParts) == 1 && "exit" == strings.ToLower(requestParts[0]) {
		fmt.Println("Exiting application.")
		return true
	}

	if len(requestParts) == 3 {
		processRequest(requestParts)
		return false
	}

	showErrorMessage("Invalid request format!")

	return false
}

func processRequest(requestParts []string) {

	switch requestParts[0] {
	case "newanimal":
		createNewAnimal(requestParts[1], requestParts[2])
	case "query":
		queryAnimal(requestParts[1], requestParts[2])
	default:
		showErrorMessage("Invalid request type!")
	}

}

func createNewAnimal(name, animalType string) {

	animal, valid := getAnimalType(animalType)
	if valid {

		animals[name] = animal
		fmt.Println("Created it!")

	}

}

func getAnimalType(animalType string) (Animal, bool) {
	switch animalType {
	case "cow":
		return Cow{"grass", "walk", "moo"}, true
	case "bird":
		return Bird{"worms", "fly", "peep"}, true
	case "snake":
		return Snake{"mice", "slither", "hsss"}, true
	}

	showErrorMessage("Invalid animal type!")

	return nil, false
}

func queryAnimal(name, information string) {

	animal, ok := animals[name]
	if ok {

		switch information {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			showErrorMessage(fmt.Sprintf("Information %s not valid!", information))
		}

	} else {

		showErrorMessage(fmt.Sprintf("Animal with name %s not found!", name))

	}

}

func showErrorMessage(message string) {

	fmt.Println()
	fmt.Println(message)
	fmt.Println()

}

func printCurrentAnimals() {

	fmt.Println()

	fmt.Println("Current animals:")
	for k, v := range animals {
		fmt.Printf("Name: %s, Type: %s\n", k, v.Type())
	}

	fmt.Println()

}
