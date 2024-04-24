package main

func main() {
	suites := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	faces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}
	for _, s := range suites {
		for _, f := range faces {
			cards = append(cards, f+" of "+s)
		}
	}

	cards.print()
}
