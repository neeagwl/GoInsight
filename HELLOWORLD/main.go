package main

import "fmt"

func main() {

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	// cards.print()

	// deal1, deal2 := deal(cards, 20)
	// deal1.print()
	// deal2.print()

	// fmt.Println(cards.toString())

	// cards := newDeckFromFile("my_cards")
	// cards.shuffle()
	// cards.print()

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {

		if num % 2 ==0 {
			fmt.Println("Even")
		}else{
			fmt.Println("Odd")
		}

	}

}
