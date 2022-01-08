package main

import (
	//"fmt"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

// Table represents the rendering of a game in progress
type TTable struct {
	widget.BaseWidget

	game *TGame
}

// CreateRenderer gets the widget renderer for this table - internal use only
func (t *TTable) CreateRenderer() fyne.WidgetRenderer {
	//fmt.Println("CreateRenderer()")
	return newRender(t)
}

// find card from an image, easier than keeping them in sync
func (t *TTable) cardFind(pos *canvas.Image) *TCard {
	//fmt.Println("cardFind()")
	deck := NewSortedDeck()
	for i, face := range deck.Cards {
		if face.Face() == pos.Resource {
			card := NewCard((i % 13), Suit(math.Floor(float64(i)/13)))
			card.FaceUp = true
			return card
		}
	}
	return nil
}

func (t *TTable) cardTapped(card *TCard) {
	//fmt.Println("cardTapped()")
	if t.game.Flag1 && !t.game.Flag2 { // State 2
		if card.Selected == false {
			card.FaceUp = false
			card.Selected = true
		} else {
			card.FaceUp = true
			card.Selected = false
		}
	}

	t.Refresh()
}

// Tapped is called when the user taps the table widget
func (t *TTable) Tapped(event *fyne.PointEvent) {
	//fmt.Println("Tapped()")
	render := test.WidgetRenderer(t).(*TRender)
	// Tapped deck
	if withinCardBounds(render.fdeck, event.Position) {

		if !t.game.Flag1 && !t.game.Flag2 && t.game.bets > 0 { // State 0
			t.game.DrawFive()
			t.game.CheckPrizes()
			render.Refresh()
			return
		}

	} else if withinCardBounds(render.fcard5, event.Position) {
		t.cardTapped(t.game.Card5)
		return
	} else if withinCardBounds(render.fcard4, event.Position) {
		t.cardTapped(t.game.Card4)
		return
	} else if withinCardBounds(render.fcard3, event.Position) {
		t.cardTapped(t.game.Card3)
		return
	} else if withinCardBounds(render.fcard2, event.Position) {
		t.cardTapped(t.game.Card2)
		return
	} else if withinCardBounds(render.fcard1, event.Position) {
		t.cardTapped(t.game.Card1)
		return
	} else {
		// Tapped elsewhere
		t.Refresh()
	}
}

// NewTable creates a new table widget for the specified game
func NewTable(g *TGame) *TTable {
	//fmt.Println("NewTable()")
	t := &TTable{game: g}

	t.game.Panel = t.newPanel()

	t.game.ButtonSet = t.newButtonSet()

	t.ExtendBaseWidget(t)
	return t
}
