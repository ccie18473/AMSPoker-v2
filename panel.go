package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TPanel struct {
	fprize    [9]*widget.Label
	fvalue    [9]*widget.Label
	fbets     *widget.Label
	fwins     *widget.Label
	fhelp     *widget.Label
	fcredits  *widget.Label
	container *fyne.Container
}

const (
	msg1 string = "Bet then tap the Deck Card"
	msg2 string = "Select those you don't want,\n then tap Retry"
	msg3 string = "Good Luck"
	msg4 string = "You won, select or credit"
	msg5 string = "You won, double or credit"
	msg6 string = "No more cards in the deck"
)

func (t *TTable) newPanel() *TPanel {
	//fmt.Println("newPanel()")
	p := &TPanel{}

	p.fprize[8] = widget.NewLabelWithStyle("Royal Flush", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[7] = widget.NewLabelWithStyle("Straight Flush", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[6] = widget.NewLabelWithStyle("4 of a Kind", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[5] = widget.NewLabelWithStyle("Full House", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[4] = widget.NewLabelWithStyle("Flush", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[3] = widget.NewLabelWithStyle("Straight", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[2] = widget.NewLabelWithStyle("3 of a Kind", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[1] = widget.NewLabelWithStyle("Two Pair", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fprize[0] = widget.NewLabelWithStyle("Pair of Aces", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})

	p.fvalue[8] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[7] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[6] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[5] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[4] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[3] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[2] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[1] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fvalue[0] = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})

	fprizes := container.NewGridWithColumns(2, p.fprize[8], p.fvalue[8],
		p.fprize[7], p.fvalue[7],
		p.fprize[6], p.fvalue[6],
		p.fprize[5], p.fvalue[5],
		p.fprize[4], p.fvalue[4],
		p.fprize[3], p.fvalue[3],
		p.fprize[2], p.fvalue[2],
		p.fprize[1], p.fvalue[1],
		p.fprize[0], p.fvalue[0],
	)

	p.fwins = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fcredits = widget.NewLabelWithStyle("100", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fbets = widget.NewLabelWithStyle("0", fyne.TextAlignLeading, fyne.TextStyle{Bold: false})
	p.fhelp = widget.NewLabelWithStyle(msg1, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})

	fstatus := container.NewGridWithColumns(2, widget.NewLabelWithStyle("Wins: ", fyne.TextAlignLeading, fyne.TextStyle{Bold: false}),
		p.fwins,
		widget.NewLabelWithStyle("Credits: ", fyne.TextAlignLeading, fyne.TextStyle{Bold: false}),
		(p.fcredits),
		widget.NewLabelWithStyle("Bets: ", fyne.TextAlignLeading, fyne.TextStyle{Bold: false}),
		(p.fbets),
	)

	fstatus = container.NewVBox(fstatus, p.fhelp)

	p.container = container.NewBorder(nil, nil, fprizes, nil, fstatus)

	return p
}
