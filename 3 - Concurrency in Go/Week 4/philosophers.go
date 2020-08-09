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

	var channelToHost = []chan string{
		make(chan string, 2),
		make(chan string, 2),
		make(chan string, 2),
		make(chan string, 2),
		make(chan string, 2),
	}

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
		go philosophers[i].eat(channelsToPhilosophers[i], channelToHost[i])
	}

	wg.Wait()
	fmt.Println("Philosophers finished the meal.")

}

func (p Philosopher) eat(in chan string, out chan string) {

	for i := 0; i < 3; i++ {

		// request authorization to eat
		fmt.Printf("Philosopher %d requesting authorization.\n", p.number)
		out <- "START"

		fmt.Printf("Philosopher %d waiting for authorization.\n", p.number)
		received := <-in

		fmt.Printf("Philosopher %d got authorization %s.\n", p.number, received)

		// take chop sticks at random order
		o := rand.Intn(2)
		if o == 0 {
			p.leftChopStick.Lock()
			p.rightChopStick.Lock()
		} else {
			p.rightChopStick.Lock()
			p.leftChopStick.Lock()
		}

		fmt.Printf("Philosopher %d is eating meal %d.\n", p.number, i)

		duration := rand.Intn(2000)
		time.Sleep(time.Duration(duration) * time.Millisecond)

		fmt.Printf("Philosopher %d finished eating meal %d.\n", p.number, i)

		p.leftChopStick.Unlock()
		p.rightChopStick.Unlock()

		// releasing authorization
		out <- "FINISH"

	}

	fmt.Printf("Philosopher %d informing Host that finished all meals.", p.number)
	out <- "COMPLETED"

	wg.Done()

	fmt.Printf("Philosopher %d completed.", p.number)

}

func host(in, out []chan string) {

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

		p, message := receiveCommunication(in)

		fmt.Printf("Host received communication %s from %d.\n", message, p)

		if eating[p] && message == "FINISH" {

			registerPhilosopherFinishedMeal(p, eating, chopSticksInUse, &eatingCount)

			pending = processPendingPhilosophers(pending, eating, chopSticksInUse, &eatingCount, out)

		} else if message == "START" && !eating[p] {

			if !checkIfPhilosopherCanEat(p, eating, chopSticksInUse, &eatingCount, out[p]) {
				pending = append(pending, p)
			}

		} else if message == "COMPLETED" {
			fmt.Printf("Host received confirmation that Philosopher %d finished all meals.\n", p)
			allFinished++
		}

		fmt.Printf("Eating Count: %d\n", eatingCount)

	}

	wg.Done()

}

func receiveCommunication(in []chan string) (int, string) {
	select {
	case message := <-in[0]:
		return 0, message
	case message := <-in[1]:
		return 1, message
	case message := <-in[2]:
		return 2, message
	case message := <-in[3]:
		return 3, message
	case message := <-in[4]:
		return 4, message
	}
}

func registerPhilosopherFinishedMeal(p int, eating []bool, chopSticksInUse []bool, eatingCount *int) {

	fmt.Printf("Host unregister Philosopher %d from eating.\n", p)

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

	fmt.Printf("Philosopher %d needs to wait.\n", p)
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
	fmt.Printf("Host sending confirmation to Philosopher %d.\n", p)
	out <- "OK"
	fmt.Printf("Host confirmed Philosopher %d.\n", p)

}
