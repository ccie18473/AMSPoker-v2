package main

import (
	//"fmt"
	"math/rand"
	"time"
)

// Deck is standard playing card collection, it contains up to 52 unique cards.
type TDeck struct {
	Cards []*TCard
}

// Shuffle reorganises the cards in the deck to a random order
func (d *TDeck) Shuffle() {
	//fmt.Println("Shuffle()")
	d.ShuffleFromSeed(time.Now().UnixNano())
}

// ShuffleFromSeed reorganises the cards in the deck to a random order using
// the specified seed for rand.Seed().
func (d *TDeck) ShuffleFromSeed(seed int64) {
	//fmt.Println("ShuffleFromSeed()")
	rand.Seed(seed)
	for c := 0; c < len(d.Cards); c++ {
		swap := rand.Intn(len(d.Cards))
		if swap != c {
			d.Cards[swap], d.Cards[c] = d.Cards[c], d.Cards[swap]
		}
	}
}

// Push adds the specified card to the top of the deck
func (d *TDeck) Push(card *TCard) {
	//fmt.Println("Push()")
	d.Cards = append(d.Cards, card)
}

// Pop removes the top card from the deck and returns it
func (d *TDeck) Pop() *TCard {
	//fmt.Println("Pop()")
	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card
}

// NewSortedDeck returns a standard deck in sorted order - starting with Ace of Clubs, ending with King of Spades.
func NewSortedDeck() *TDeck {
	//fmt.Println("NewSortedDeck()")
	d := &TDeck{}

	c := 0
	suit := SuitClubs
	for i := 0; i < 4; i++ {
		for value := 0; value <= ValueKing; value++ {
			d.Cards = append(d.Cards, NewCard(value, suit))
			c++
		}
		suit++
	}

	return d
}

// NewShuffledDeck returns a 52 card deck in random order.
func NewShuffledDeck() *TDeck {
	//fmt.Println("NewShuffledDeck()")
	d := NewSortedDeck()
	d.Shuffle()

	return d
}

// NewShuffledDeckFromSeed returns a 52 card deck in random order.
// The randomness is seeded using the seed parameter.
func NewShuffledDeckFromSeed(seed int64) *TDeck {
	//fmt.Println("NewShuffledDeckFromSeed()")
	d := NewSortedDeck()
	d.ShuffleFromSeed(seed)

	return d
}
