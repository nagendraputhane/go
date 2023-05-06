package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	//variable name string = "string"

	cards := "Ace of Spades"
	//same as var cards string
	//:= only when declaring new variable
	card = "Five of Diamonds"
	fmt.Println(card, cards)
}
