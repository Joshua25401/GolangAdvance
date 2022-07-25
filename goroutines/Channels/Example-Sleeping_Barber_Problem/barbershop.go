package main

import (
	"fmt"
	"time"
)

type Barbershop struct {
	ShopCapacity       int
	HairCutDuration    time.Duration
	NumberOfBarbers    int
	BarbersDoneChannel chan bool
	ClientChannel      chan string
	Open               bool
}

func (b *Barbershop) addBarber(barber string) {
	// Increment attribute
	b.NumberOfBarbers++

	go func() {
		isSleeping := false
		fmt.Printf("%s goes to waiting rooms to check for clients!\n", barber)

		// Do hair cut here!
		for {
			// If there are no clients, the barbers goes to sleep
			if len(b.ClientChannel) == 0 {
				fmt.Printf("There's nothing to do, so %s takes a nap!\n", barber)
				isSleeping = true
			}

			client, shopOpen := <-b.ClientChannel

			if shopOpen {
				if isSleeping {
					fmt.Printf("%s wakes %s up!\n", client, barber)
					isSleeping = false
				}

				// Cut the clients hair
				b.cutHair(barber, client)
			} else {
				// Shop closed -> Barber goes to home and close this goroutine
				b.sendBarberToHome(barber)
				return
			}
		}
	}()
}

func (b *Barbershop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair!\n", barber, client)
	time.Sleep(b.HairCutDuration)
	fmt.Printf("%s is finished to cutting %s's hairs!\n", barber, client)
}

func (b *Barbershop) sendBarberToHome(barber string) {
	fmt.Printf("%s is going home!\n", barber)
	b.BarbersDoneChannel <- true
}

func (b *Barbershop) closeShopForThisDay() {
	fmt.Printf("Closing shop for this day!\n")
	close(b.ClientChannel)
	b.Open = false

	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChannel
	}
	close(b.BarbersDoneChannel)
	fmt.Println("==================================================")
	fmt.Printf("Barbershop close for the day! Everyone goes home!\n")
}

func (b *Barbershop) addClient(client string) {
	fmt.Printf("===> %s arrives!\n", client)

	if b.Open {
		select {
		case b.ClientChannel <- client:
			fmt.Printf("%s takes a seat in waiting room!\n", client)

		default:
			fmt.Printf("Waiting room is full! so, %s leaves!\n", client)
		}
	} else {
		fmt.Printf("The Barbershop is already closed! %s leaves!\n", client)
	}
}
