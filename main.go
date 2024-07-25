package main

import "fmt"

func main(){
	var card string = "Ace of Spades"
	fmt.Println(card)
	// Another method of initialization
	card1 := "Five of Diamonds"
	// Since the variable is already intialized, so we are not using colon(:), just assigning =
	card1 = "Six of Hearts" 
	fmt.Println(card1) 

	card2 := newCard()
	fmt.Println(card2)
}

func newCard() string{
	return "Ace of Clubs"
}