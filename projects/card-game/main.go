package main

import "fmt"

func main() {
	deck1 := newDeck()
	//deck1.print()

	hand, remaining := deal(deck1, 5)
	hand.print()
	fmt.Println("Remaining")
	remaining.print()
}
