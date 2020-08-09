package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ChopStick Represents a Philosopher ChopStick to eat a Meal
type ChopStick struct{ sync.Mutex }

// Philosopher represents a philosopher at dinning table
type Philosopher struct {
	number                        int
	leftChopStick, rightChopStick *ChopStick
}

// PhilosopherAction represents an action of Philosopher in table
type PhilosopherAction struct {
	number  int
	message string
}

var wg sync.WaitGroup

func main() {

	chopSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopSticks[i], chopSticks[(i+1)%5]}
	}

	var channelToHost chan PhilosopherAction = make(chan PhilosopherAction, 10)

	var channelsToPhilosophers = []chan string{
		make(chan string, 1),
		make(chan string, 1),
		make(chan string, 1),
		make(chan string, 1),
		make(chan string, 1),
	}

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Philosophers started the meal.")
	wg.Add(6)

	go host(channelToHost, channelsToPhilosophers)

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(channelsToPhilosophers[i], channelToHost)
	}

	wg.Wait()
	fmt.Println("Philosophers finished the meal.")

}

func (p Philosopher) eat(in chan string, out chan PhilosopherAction) {

	for i := 0; i < 3; i++ {

		// request authorization to eat
		fmt.Printf("PHILOSOPHER %d: requesting authorization.\n", p.number)
		out <- PhilosopherAction{p.number, "START"}

		fmt.Printf("PHILOSOPHER %d: waiting for authorization.\n", p.number)
		received := <-in

		fmt.Printf("PHILOSOPHER %d: got authorization %s.\n", p.number, received)

		// take chop sticks at random order
		o := rand.Intn(2)
		if o == 0 {
			p.leftChopStick.Lock()
			p.rightChopStick.Lock()
		} else {
			p.rightChopStick.Lock()
			p.leftChopStick.Lock()
		}

		fmt.Printf("PHILOSOPHER %d: is eating meal %d.\n", p.number, i)

		duration := rand.Intn(2000)
		time.Sleep(time.Duration(duration) * time.Millisecond)

		fmt.Printf("PHILOSOPHER %d: finished eating meal %d.\n", p.number, i)

		p.leftChopStick.Unlock()
		p.rightChopStick.Unlock()

		// releasing authorization
		out <- PhilosopherAction{p.number, "FINISH"}

	}

	fmt.Printf("PHILOSOPHER %d: informing Host that finished all meals.\n", p.number)
	out <- PhilosopherAction{p.number, "COMPLETED"}

	wg.Done()

	fmt.Printf("PHILOSOPHER %d: completed.\n", p.number)

}

func host(in chan PhilosopherAction, out []chan string) {

	chopSticksInUse := make([]bool, 5)
	for i := range chopSticksInUse {
		chopSticksInUse[i] = false
	}

	eating := make([]bool, 5)
	for i := range eating {
		eating[i] = false
	}

	var pending []int = make([]int, 0)

	eatingCount := 0
	allFinished := 0

	fmt.Println("Host started listening for Philosophers authorization.")
	for allFinished < 5 {

		action := <-in

		fmt.Printf("HOST: received communication %s from %d.\n", action.message, action.number)

		if eating[action.number] && action.message == "FINISH" {

			registerPhilosopherFinishedMeal(action.number, eating, chopSticksInUse, &eatingCount)

			pending = processPendingPhilosophers(pending, eating, chopSticksInUse, &eatingCount, out)

		} else if action.message == "START" && !eating[action.number] {

			if !checkIfPhilosopherCanEat(action.number, eating, chopSticksInUse, &eatingCount, out[action.number]) {
				pending = append(pending, action.number)
			}

		} else if action.message == "COMPLETED" {
			fmt.Printf("HOST: received confirmation that Philosopher %d finished all meals.\n", action.number)
			allFinished++
		}

		fmt.Printf("HOST: eating count %d\n", eatingCount)

	}

	wg.Done()

}

func registerPhilosopherFinishedMeal(p int, eating []bool, chopSticksInUse []bool, eatingCount *int) {

	fmt.Printf("HOST: unregister Philosopher %d from eating.\n", p)

	eating[p] = false
	chopSticksInUse[p] = false
	chopSticksInUse[(p+1)%5] = false
	*eatingCount--
}

func checkIfPhilosopherCanEat(p int, eating []bool, chopSticksInUse []bool, eatingCount *int, out chan string) bool {

	leftChopStick := p
	rightChopStick := (p + 1) % 5

	if *eatingCount < 2 && !eating[p] && !chopSticksInUse[leftChopStick] && !chopSticksInUse[rightChopStick] {
		allowPhilosopherToEat(p, eating, chopSticksInUse, eatingCount, out)
		return true
	}

	fmt.Printf("HOST: Philosopher %d needs to wait.\n", p)
	return false

}

func processPendingPhilosophers(pending []int, eating, chopSticksInUse []bool, eatingCount *int, out []chan string) []int {

	for i, p := range pending {

		leftChopStick := p
		rightChopStick := (p + 1) % 5

		if !chopSticksInUse[leftChopStick] && !chopSticksInUse[rightChopStick] {
			allowPhilosopherToEat(p, eating, chopSticksInUse, eatingCount, out[p])
			return removePending(pending, i)
		}

	}

	return pending
}

func removePending(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func allowPhilosopherToEat(p int, eating, chopSticksInUse []bool, eatingCount *int, out chan string) {

	eating[p] = true
	chopSticksInUse[p] = true
	chopSticksInUse[(p+1)%5] = true
	*eatingCount++
	fmt.Printf("HOST: sending confirmation to Philosopher %d.\n", p)
	out <- "OK"
	fmt.Printf("HOST: confirmed Philosopher %d.\n", p)

}
