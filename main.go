package main

import "fmt"

func main(){
	var card string = "A♠️"
	fmt.Println(card)
	// Another method of initialization
	card1 := "5♦️"
	// Since the variable is already intialized, so we are not using colon(:), just assigning =
	card1 = "6❤️" 
	fmt.Println(card1) 

	card2 := newCard()
	fmt.Println(card2)

	cards:= [] string {}
	cards = append(cards, "5♦️", card, card1, newCard() )
	fmt.Println(cards)
	for i, crd:= range cards{
		fmt.Println(i, crd)
	}
}

func newCard() string{
	return "A♣️"
}