package main

import "fmt"

// Entry main function
func main(){
	cards := newDeck()
	rightSide, leftSide := deal(cards, 5)
	
	rightSide.print()
	fmt.Println()
	leftSide.print()
}