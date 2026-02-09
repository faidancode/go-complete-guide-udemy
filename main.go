package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Four of Clubs")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
