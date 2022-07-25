package main

import (
	"fmt"
	"sync"
)

/*
	Goroutine -> Very lightweight process

	How to use Goroutine :
		1. Write 'go' keyword to create goroutine.

	NOTE:
		- Main function in go is also 1 Goroutine
		- If program die or end quickly the goroutine may not come back from background
*/

func printName(s string) {
	fmt.Println(s)
}

/*
	Here we add more parameters such as wg
	wg is an object of sync.WaitGroup

	Purpose :
		- to tell the WaitGroup object that the process is done
		  and wait for the other process or goroutines.
*/
func printName2(name string, wg *sync.WaitGroup) {
	defer wg.Done() // Tell other goroutines that this function done
	fmt.Println(name)
}

func main() {
	// Un-comment this for example of the goroutine not come back
	//go printName("Joshua Ryandafres Pangaribuan")
	//printName("Irwan Rivandy Siagian")

	//// Solve the problem above!
	//// NOTE : is not a good solution. just for example!
	//go printName("Joshua Ryandafres Pangaribuan")
	//// Adding more time to 'wait' the goroutine
	//time.Sleep(1 * time.Second)
	//printName("Irwan Rivandy Siagian")

	names := []string{
		"Joshua Ryandafres Pangaribuan",
		"Irwan Rivandy Siagian",
		"Wilson",
		"Rotua Lumbangaol",
		"Yogie Sinaga",
	}

	var wg sync.WaitGroup
	wg.Add(len(names))

	for indeks, name := range names {
		go printName2(fmt.Sprintf("%d -> %s", indeks, name), &wg)
	}

	// Really bad idea to implement time.Sleep when work with goroutine
	// Use WaitGroup instead!
	//time.Sleep(1 * time.Second)

	// Instead of using time.Sleep
	// we use WaitGroup here!
	// wg.Wait() -> WaitGroup function to wait all gouroutines done executing
	wg.Wait()

	fmt.Println("Program Stopped!")
}
