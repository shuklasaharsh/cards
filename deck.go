package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
)


type deck []Card



func getNewDeck() deck {
	var cards deck
	for i:=1; i<=4; i++ {
		for j:=1; j<=13; j++ {
			cards = append(cards, getNewCard(j, i))
		}
	}
	return cards
}

func (d deck) print() {
	for _, card:= range d {
		fmt.Println(card.Suit, card.FaceValue, card.Text)
	}
}

func (d deck) shuffle() {

}

func deal(handSize int, d deck) (deck, deck){
	h:= d[:handSize]
	d=d[handSize:]
	return h, d
}

func (d deck) saveToFile(fileName string) error {
	jsonData, err := json.Marshal(d)
	if err != nil {
		panic(err)
		return err
	}

	modePerm := fs.ModeAppend
	err = ioutil.WriteFile(fileName, jsonData, modePerm)
	if err != nil {
		panic(err)
		return err
	}

	return nil
}

func (d deck) saveHands() error {
	return nil
}

func readFromFile() (deck, error) {
	var d deck
	data, err := ioutil.ReadFile("./state.json")
	if err != nil {
		panic(err)
		return nil, err
	}
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}
	fmt.Println(d)
	return d, err
}