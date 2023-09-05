package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck' which is  a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
func (d deck) toString() string {

	return strings.Join([]string(d), ",")

}
func newDeckFromFile(filepath string) deck {
	bs, error := ioutil.ReadFile(filepath)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for index := range d {
		position := r.Intn(len(d) - 1)
		// temp := d[index]
		d[index], d[position] = d[position], d[index]
		// d[position] = temp
	}
}

// receiver on a function using the parameter first
func (d deck) print() {
	for ind, ele := range d {
		fmt.Println(ind, ele)
	}
}
