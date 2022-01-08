package main

import (
	//"fmt"
	"log"

	"fyne.io/fyne/v2"

	"poker/faces"
)

// Suit encodes one of the four possible suits for a playing card
type Suit int

// SuitColor represents the red/black of a suit
type SuitColor int

const (
	// SuitClubs is the "Clubs" playing card suit
	SuitClubs Suit = iota
	// SuitDiamonds is the "Diamonds" playing card suit
	SuitDiamonds
	// SuitHearts is the "Hearts" playing card suit
	SuitHearts
	// SuitSpades is the "Spades" playing card suit
	SuitSpades

	// SuitColorBlack is returned from Color() if the suit is Clubs or Spades
	SuitColorBlack SuitColor = iota
	// SuitColorRed is returned from Color() if the suit is Diamonds or Hearts
	SuitColorRed
)

const (
	// ValueJack is a convenience for the card 1 higher than 10
	ValueJack = 10
	// ValueQueen is the value for a queen face card
	ValueQueen = 11
	// ValueKing is the value for a king face card
	ValueKing = 12
)

// Card is a single playing card, it has a face value and a suit associated with it.
type TCard struct {
	Value int
	Suit  Suit

	FaceUp   bool
	Selected bool
}

// Face returns a resource that can be used to render the associated card
func (c *TCard) Face() fyne.Resource {
	//fmt.Println("Face()")
	return faces.ForCard(c.Value, int(c.Suit))
}

// TurnFaceUp sets the FaceUp field to true - so the card value can be seen
func (c *TCard) TurnFaceUp() {
	//fmt.Println("TurnFaceUp()")
	c.FaceUp = true
}

// TurnFaceDown sets the FaceUp field to false - so the card should be hidden
func (c *TCard) TurnFaceDown() {
	//fmt.Println("TurnFaceDown()")
	c.FaceUp = false
}

// Color returns the red or black color of the card suit
func (c *TCard) Color() SuitColor {
	//fmt.Println("Color()")
	if c.Suit == SuitClubs || c.Suit == SuitSpades {
		return SuitColorBlack
	}

	return SuitColorRed
}

// NewCard returns a new card instance with the specified suit and value (1 based for Ace, 2 is 2 and so on).
func NewCard(value int, suit Suit) *TCard {
	//fmt.Println("NewCard()")
	if value < 0 || value > 12 {
		log.Fatal("Invalid card face value")
	}

	return &TCard{Value: value, Suit: suit}
}
