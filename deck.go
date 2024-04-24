package main

import "fmt"

type deck []string

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	suites := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	faces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}
	for _, s := range suites {
		for _, f := range faces {
			cards = append(cards, f+" of "+s)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
