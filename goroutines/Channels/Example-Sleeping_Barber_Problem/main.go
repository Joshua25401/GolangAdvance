package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Variables
var seatCapacity = 10
var customerArrivalRate = 100
var haircutDuration = time.Duration(rand.Intn(2000)+1) * time.Millisecond
var openTime = 10 * time.Second

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Print welcome message
	fmt.Println("SLEEPING BARBER PROBLEM")
	fmt.Println("=======================")

	// Create channel here if we need any
	clientChannel := make(chan string, seatCapacity) // Buffered channel
	doneChannel := make(chan bool)                   // Whenever done send bool value to this channel

	// Create the barbershop
	shop := Barbershop{
		ShopCapacity:       seatCapacity,
		HairCutDuration:    haircutDuration,
		NumberOfBarbers:    0,
		BarbersDoneChannel: doneChannel,
		ClientChannel:      clientChannel,
		Open:               true,
	}
	fmt.Println("Barbershop is open for the day!")

	// Add barbers
	shop.addBarber("Joshua")
	shop.addBarber("Irwan")
	shop.addBarber("Rotua")
	shop.addBarber("Yogie")

	// Start barbershop as a goroutine (background)
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(openTime)
		shopClosing <- true
		shop.closeShopForThisDay()
		closed <- true
	}()

	// Add clients / customers
	i := 1
	go func() {
		for {
			// Get a random number with average arrival rate
			randMil := rand.Int() % (2 * customerArrivalRate)
			select {
			case <-shopClosing:
				return

			case <-time.After(time.Millisecond * time.Duration(randMil)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// Block until the barbershop closed (keep the app running)
	<-closed
}
