package main

import "fmt"

func main() {
	//array - fixed. Same datatype
	//slice - array that can grown and shrink
	cards := []string{"Ace of Diamonds", newCard()}
	//var = []datatype{}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		//for index, current item := range slice {}
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
