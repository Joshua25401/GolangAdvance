package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

// Pizza Constraints
const NumberOfPizzas = 10

/*
	Here we have 3 Variables that we use to keep track on :
	- How many pizzas was successfuly created (pizzasMade)
	- How many pizzas was failed (pizzasFailed)
	- Total of pizzas that created either success or failed (total)

	Default is 0 (int)
*/

var pizzasMade, pizzasFailed, total int

// Here we define Producer struct
// That create the pizzas
type Producer struct {
	dataPizza chan PizzaOrder
	quit      chan chan error
}

/*
	Here we define PizzaOrder struct
	It will hold 3 data such as :
	1. Number of Pizza (pizzaNumber)
	2. Message (We use to hold notification if pizza success of failed)
	3. Success (We use to track pizza are success or failed)
*/
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	//Increment pizzaNumber
	// Indicate that we have pizza that we working on now
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1 // So, the delay never be 0 or below that
		fmt.Printf("Received order number #%v!\n", pizzaNumber)

		// Try to make pizza
		orderCondition := rand.Intn(12) + 1
		msg := ""
		success := false
		// Check the orderCondition
		// if less than 5 -> Assume that pizza failed to make
		// if greater than 5 -> Assume that pizza success to make
		if orderCondition <= 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		// Increment the total
		total++

		fmt.Printf("Making pizza number #%v. It'll take %v seconds...\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		// Create the message either it successful pizza or failed pizza
		if orderCondition <= 2 {
			msg = fmt.Sprintf("FAILED\t: We ran out of ingredients for pizza %v!\n", pizzaNumber)
		} else if orderCondition <= 5 {
			msg = fmt.Sprintf("FAILED\t: The cook quit when making pizza number %v!\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("SUCCESS\t: Pizza order number %v is ready!\n", pizzaNumber)
		}

		pizzaOrder := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &pizzaOrder
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzaria(pizzaProducer *Producer) {
	// Keep track of which pizza we are making
	var pizzaIndeks = 0

	// Run infinite loop until receive a quit notification
	// Try to make pizzas
	for {
		// Process to make the pizza
		currentPizza := makePizza(pizzaIndeks)
		if currentPizza != nil {
			pizzaIndeks = currentPizza.pizzaNumber

			select {
			case pizzaProducer.dataPizza <- *currentPizza:
			case quitChan := <-pizzaProducer.quit:
				close(pizzaProducer.dataPizza)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano()) // Prevent we get same result of rand number

	// Print messages
	color.Cyan("The Pizzaria is open for business!\n")
	color.Cyan("==================================\n")

	// Create a Pizza producer
	pizzaProducer := &Producer{
		dataPizza: make(chan PizzaOrder),
		quit:      make(chan chan error),
	}

	// Running producer in the background
	go pizzaria(pizzaProducer)

	// Create and run consumer
	for i := range pizzaProducer.dataPizza {
		if i.pizzaNumber <= NumberOfPizzas {
			// Try to order a pizza
			if i.success {
				color.Green(i.message)
				color.Green("Order of #%v is out for delivery!\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!\n")
			}
		} else {
			fmt.Println("Done making pizza...")
			err := pizzaProducer.Close()
			if err != nil {
				fmt.Println("ERROR\t: Closing channel", err)
			}
		}
	}

	// Print out the ending message
	fmt.Println("------------------")
	fmt.Println("Done for this day!")
	fmt.Printf("We made %v pizzas, but failed to make %v, with %v attempts in total\n", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		fmt.Println("It's an awful day...")

	case pizzasFailed >= 6:
		fmt.Println("It's not a very good day...")

	case pizzasFailed >= 2:
		fmt.Println("It was a great day!!!")
	}
}
