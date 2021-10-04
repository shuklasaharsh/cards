package main

import "fmt"

func main() {
	deck := getNewDeck()
	deck.print()
	fmt.Println()
	hand, deck:=deal(5, deck)
	hand.saveToFile("./hand.json")
	deck.saveToFile("./state.json")
}