package main

func main() {
	// cards1 := deck{"ace of Diamonds", newCard()}
	// cards2 := append(cards1, "Five of Hearts")
	// fmt.Println(cards1)
	// fmt.Println(cards2)
	// for i, card := range cards2 {
	// 	fmt.Println(i, card)
	// }
	// var i2 int
	// var card2 string
	// for i2, card2 = range cards1 {
	// 	fmt.Println(i2, card2)
	// }
	//cards1.print()
	// var h2 deck
	// cards := newDeck()
	// cards.print()
	// h1, h2 := deal(cards, 6)
	// fmt.Println("hand 1")
	// h1.print()
	// fmt.Println("hand 2")
	// h2.print()

	// sliceOfBytes := []byte("hello")
	// fmt.Println(sliceOfBytes)
	// fmt.Println(h1.toString())

	// h1.saveToFile("acd.txt")
	// fmt.Println("read file")
	// h3 := newDeckFromFile("acd.txt")
	// h3.print()

	cards5 := newDeck()
	cards5.shuffle()
	cards5.print()

}

// func newCard() string {
// 	return "Five of Diamonds"
// }
