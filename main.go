package main

import "fmt"

func main() {
	suites := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	faces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := []string{}
	for _, s := range suites {
		for _, f := range faces {
			cards = append(cards, f+" of "+s)
		}
	}
	for _, d := range cards {
		fmt.Println(d)
	}
}
