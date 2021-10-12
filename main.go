package main

import "fmt"

func main() {
	deck := getNewDeck()
	deck.print()
	fmt.Println()
	hand, deck :=deal(5, deck)
	err := hand.saveToFile("./hand.json")
	if err != nil {
		panic(err)
	}
	err = deck.saveToFile("./state.json")
	if err != nil {
		panic(err)
	}
}