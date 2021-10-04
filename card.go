package main

type Card struct {
	Text      string `json:"text"`
	FaceValue int `json:"faceValue"`
	Suit      int `json:"suit"`
}

func getNewCard(faceValue, suit int) Card {
	var value, suits string
	switch faceValue {
	case 1:
		value = "Ace"
	case 2:
		value = "Two"
	case 3:
		value = "Three"
	case 4:
		value = "Four"
	case 5:
		value = "Five"
	case 6:
		value = "Six"
	case 7:
		value = "Seven"
	case 8:
		value = "Eight"
	case 9:
		value = "Nine"
	case 10:
		value = "Ten"
	case 11:
		value = "Jack"
	case 12:
		value = "Queen"
	case 13:
		value = "King"
	default:
		value = "Joker"
	}

	switch suit {
	case 1:
		suits = "Spades"
	case 2:
		suits = "Diamonds"
	case 3:
		suits = "Clubs"
	case 4:
		suits = "Hearts"
	default:
		suits = ""
	}

	return Card{
		value + " of " + suits,
		faceValue,
		suit,
	}
}
