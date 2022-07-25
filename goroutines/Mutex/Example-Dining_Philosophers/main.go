package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"sync"
	"time"
)

// Max Philoshopers eating
const maxEat = 3

// Mutex
var orderMutex sync.Mutex

// WaitGroup
var wg sync.WaitGroup

// SleepTime
var sleepTime = 1 * time.Second

// EatTime
var eatTime = 2 * time.Second

// ThinkTime
var thinkTime = 1 * time.Second

// Variables
var philosophers = []string{
	"Joshua",
	"Irwan",
	"Wilson",
	"Yogie",
	"Rotua",
}

// OrderFinished
var orderFinished []string

// diningTabel func
func diningTable(philosopher string, rightFork, leftFork *sync.Mutex) {
	defer wg.Done()

	// Print message
	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for indeks := maxEat; indeks > 0; indeks-- {
		fmt.Println(philosopher, "is hungry!")
		time.Sleep(sleepTime)
		// Lock both fort (rightFork and leftFork)
		leftFork.Lock()
		fmt.Println("\t", philosopher, "picked up the fork to his left hand!")
		rightFork.Lock()
		fmt.Println("\t", philosopher, "picked up the fork to his right hand!")

		// Print a message
		color.Green("\t" + philosopher + " has both fork and start eating!\n")
		time.Sleep(eatTime)

		// Give philosopher time to think
		fmt.Println("\t", philosopher, "start to thinking!")
		time.Sleep(thinkTime)

		// Unlock the fork
		leftFork.Unlock()
		rightFork.Unlock()
		color.Yellow("\t" + philosopher + " already eating and put both fork back!\n")
		time.Sleep(sleepTime)
	}

	// Print done message
	color.Green(philosopher + " is statisfied!\n")
	time.Sleep(sleepTime)

	color.Red(philosopher + " has left the table!")
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}

func main() {
	// Print Intro
	fmt.Println("=== Dining Philosopher Problem ===")
	fmt.Println("==================================")

	// Spawn goroutine for each philosopher
	wg.Add(len(philosophers))

	leftFork := &sync.Mutex{}

	for indeks := 0; indeks < len(philosophers); indeks++ {
		// Create a mutex for the right fork
		rightFork := &sync.Mutex{}

		// Call goroutine here
		go diningTable(philosophers[indeks], rightFork, leftFork)

		leftFork = rightFork
	}

	wg.Wait()
	fmt.Println("=== Problem Finished! ===")
	fmt.Println("Order : ", strings.Join(orderFinished, ", "))
}
