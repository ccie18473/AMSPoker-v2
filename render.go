package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	"poker/faces"
)

const minPadding = float32(4)

const minCardWidth = float32(95)
const cardRatio = 142.0 / minCardWidth



var (
	cardSize = fyne.Size{Width: minCardWidth, Height: minCardWidth * cardRatio}
	panelSize = fyne.Size{Width: 7 * minCardWidth, Height: 2 * minCardWidth * cardRatio}
	buttonsetSize = fyne.Size{Width: 7 * minCardWidth, Height: minCardWidth / 3}

	smallPad = minPadding
)

func updateCardPosition(c *canvas.Image, x, y float32) {
	//fmt.Println("updateCardPosition()")
	c.Resize(cardSize)
	c.Move(fyne.NewPos(x, y))
}

func updateButtonSetPosition(b fyne.CanvasObject, x, y float32) {
	//fmt.Println("updateButtonSetPosition()")
	b.Resize(buttonsetSize)
	b.Move(fyne.NewPos(x, y))
}

func updatePanelPosition(p fyne.CanvasObject, x, y float32) {
	//fmt.Println("updatePanelPosition()")
	p.Resize(panelSize)
	p.Move(fyne.NewPos(x, y))
}

func withinCardBounds(c *canvas.Image, pos fyne.Position) bool {
	//fmt.Println("withinCardBounds()")
	if pos.X < c.Position().X || pos.Y < c.Position().Y {
		return false
	}

	if pos.X >= c.Position().X+c.Size().Width || pos.Y >= c.Position().Y+c.Size().Height {
		return false
	}

	return true
}

func newCardObject(card *TCard) *canvas.Image {
	//fmt.Println("newCardObject()")
	if card == nil {
		return &canvas.Image{}
	}

	var face fyne.Resource
	if card.FaceUp {
		face = card.Face()
	} else {
		face = faces.ForBack()
	}
	image := &canvas.Image{Resource: face}
	image.Resize(cardSize)

	return image
}

func newCardSpace() *canvas.Image {
	//fmt.Println("newCardSpace()")
	space := faces.ForSpace()
	image := &canvas.Image{Resource: space}
	image.Resize(cardSize)

	return image
}

func newButtonSetObject(ButtonSet *TButtonSet) *fyne.Container {
	//fmt.Println("newButtonSetObject()")
	if ButtonSet == nil {
		return nil
	}
	return ButtonSet.container
}

func newPanelObject(Panel *TPanel) *fyne.Container {
	//fmt.Println("newPanelObject()")
	if Panel == nil {
		return nil
	}
	return Panel.container
}

type TRender struct {
	game *TGame

	fdeck *canvas.Image

	fcard1, fcard2, fcard3, fcard4, fcard5, fcard6                       *canvas.Image
	fbuttonset															 fyne.CanvasObject
	fpanel                                                               fyne.CanvasObject

	fobjects []fyne.CanvasObject
	table    *TTable
}

func (r *TRender) MinSize() fyne.Size {
	//fmt.Println("MinSize()")
	return fyne.NewSize(720, 540)
}

func (r *TRender) Layout(size fyne.Size) {
	//fmt.Println("Layout()")

	// Deck
	updateCardPosition(r.fdeck, smallPad, smallPad)
	// Cards
	updateCardPosition(r.fcard1, smallPad * 6 + cardSize.Width, smallPad)
	updateCardPosition(r.fcard2, smallPad * 6 + cardSize.Width * 2, smallPad)
	updateCardPosition(r.fcard3, smallPad * 6 + cardSize.Width * 3, smallPad)
	updateCardPosition(r.fcard4, smallPad * 6 + cardSize.Width * 4, smallPad)
	updateCardPosition(r.fcard5, smallPad * 6 + cardSize.Width * 5, smallPad)
	// Double Card
	updateCardPosition(r.fcard6, smallPad * 12 + cardSize.Width * 6, smallPad)
	
	// Panel
	updatePanelPosition(r.fpanel, smallPad + cardSize.Width, smallPad * 4 + cardSize.Height)
	
	// Buttons
	updateButtonSetPosition(r.fbuttonset, smallPad * 6, cardSize.Height * 3.5)
}

func (r *TRender) ApplyTheme() {
	//fmt.Println("ApplyTheme()")
	// no-op we are a custom UI
}

func (r *TRender) refreshCard(image *canvas.Image, card *TCard) {
	//fmt.Println("refreshCard()")
	image.Hidden = card == nil

	image.Resource = faces.ForSpace()
	image.Translucency = 0
	if card == nil {
		image.Resource = faces.ForSpace()
		return
	}

	if card.FaceUp {
		image.Resource = card.Face()
	} else {
		image.Resource = faces.ForSpace()
	}
}

func (r *TRender) refreshButtonSet(buttonset fyne.CanvasObject, ButtonSet *TButtonSet) {
	//fmt.Println("refreshButtonSet()")
	if buttonset == nil {
		return
	}
	buttonset = ButtonSet.container
}

func (r *TRender) refreshPanel(fpanel fyne.CanvasObject, Panel *TPanel) {
	//fmt.Println("refreshPanel()")
	if fpanel == nil {
		return
	}
	fpanel = Panel.container
}

func (r *TRender) Refresh() {
	//fmt.Println("Refresh()")
	if len(r.game.Deck.Cards) > 0 {
		r.fdeck.Resource = faces.ForBack()
	} else {
		r.fdeck.Resource = faces.ForSpace()
	}

	canvas.Refresh(r.fdeck)

	if !r.game.Flag4 { // Waiting for doubling
		r.fcard6.Resource = faces.ForBack()
		canvas.Refresh(r.fcard6)
	} else {
		r.refreshCard(r.fcard6, r.game.Card6)
	}

	r.refreshCard(r.fcard1, r.game.Card1)
	r.refreshCard(r.fcard2, r.game.Card2)
	r.refreshCard(r.fcard3, r.game.Card3)
	r.refreshCard(r.fcard4, r.game.Card4)
	r.refreshCard(r.fcard5, r.game.Card5)

	r.refreshButtonSet(r.fbuttonset, r.game.ButtonSet)

	r.refreshPanel(r.fpanel, r.game.Panel)

	canvas.Refresh(r.table)
}

func (r *TRender) Objects() []fyne.CanvasObject {
	//fmt.Println("Objects()")
	return r.fobjects
}

func (r *TRender) Destroy() {
	//fmt.Println("Destroy()")
}

func newRender(table *TTable) *TRender {
	//fmt.Println("newRender()")
	r := &TRender{}
	r.table = table
	r.game = table.game
	r.fdeck = newCardObject(nil)

	r.fcard1 = newCardObject(nil)
	r.fcard2 = newCardObject(nil)
	r.fcard3 = newCardObject(nil)
	r.fcard4 = newCardObject(nil)
	r.fcard5 = newCardObject(nil)

	r.fcard6 = newCardObject(nil)

	r.fbuttonset = newButtonSetObject(r.game.ButtonSet)

	r.fpanel = newPanelObject(r.game.Panel)

	r.fobjects = []fyne.CanvasObject{r.fdeck, r.fcard1, r.fcard2, r.fcard3,
		r.fcard4, r.fcard5, r.fcard6, r.fpanel, r.fbuttonset}

	r.Refresh()
	return r
}
