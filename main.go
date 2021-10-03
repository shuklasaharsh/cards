package main

import "fmt"



func main() {
	cards:= getNewDeck()
	for _, el := range cards {
		fmt.Println(el.text)
	}


}


