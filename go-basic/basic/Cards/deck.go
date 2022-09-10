package main

import "fmt"

type Deck []string

func newDeck() Deck{
	cards := Deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits{
		for _, value := range cardsValues{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func deal (d Deck, handSize int)(Deck, Deck){
	return d[:handSize], d[handSize:]
}

func (d Deck) print(){ // Value receiver (just copy the value of Deck not the reference)
	for i, card := range d {
		fmt.Printf("%v %v\n",i,card)
	}
}