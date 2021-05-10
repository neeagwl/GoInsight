package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string


func newDeck() deck{

    cards := deck{}

    cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
    cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", 
                            "Eight", "Nine", "Ten", "Jack", "Queen", "King" }

    for _,suit := range cardSuits{
        for _,value := range cardValues{

            cards = append(cards, value +" of " + suit)            

        }

    }                        

    return cards

}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, hand int) (deck, deck) {

    return d[:hand], d[hand:]

}

func (d deck) toString() string{

    return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error{

    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck{

    bs, err := ioutil.ReadFile(filename)

    if err!=nil{
        fmt.Println("Error: ", err)
        os.Exit(1)
    }

    s := strings.Split(string(bs), ",")

    return deck(s)

}


func (d deck) shuffle(){

    src := rand.NewSource(time.Now().UnixNano())
    r := rand.New(src)

    for i := range d{

        newPos := r.Intn(len(d))

        d[i], d[newPos] = d[newPos], d[i]
    }

}