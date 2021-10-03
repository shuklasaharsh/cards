package main

// Create a new type of deck
// This is a slice of strings
type deck []Card

func getNewDeck() deck {
	var cards deck
	for i:=1; i<=4; i++ {
		for j:=1; j<=14; j++ {
			cards = append(cards, getNewCard(j, i))
		}
	}
	return cards
}
