package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	Type() string
}

type Cow struct{ eat, move, speak string }
type Bird struct{ eat, move, speak string }
type Snake struct{ eat, move, speak string }

func (c Cow) Eat() {
	fmt.Printf("Cow eats %s\n", c.eat)
}

func (c Cow) Move() {
	fmt.Printf("Cow moves %s\n", c.move)
}

func (c Cow) Speak() {
	fmt.Printf("Cow makes noise %s\n", c.speak)
}

func (c Cow) Type() string {
	return "Cow"
}

func (b Bird) Eat() {
	fmt.Printf("Bird eats %s\n", b.eat)
}

func (b Bird) Move() {
	fmt.Printf("Bird moves %s\n", b.move)
}

func (b Bird) Speak() {
	fmt.Printf("Bird makes noise %s\n", b.speak)
}

func (b Bird) Type() string {
	return "Bird"
}

func (s Snake) Eat() {
	fmt.Printf("Snake eats %s\n", s.eat)
}

func (s Snake) Move() {
	fmt.Printf("Snake moves %s\n", s.move)
}

func (s Snake) Speak() {
	fmt.Println("Snake makes noise %s\n", s.speak)
}

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
