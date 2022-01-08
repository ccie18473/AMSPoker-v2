package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TButtonSet struct {
	button1, button2, button3, button4, button5, button6, button7    *widget.Button
	container *fyne.Container
}

const buttonWidth float32 = 200
const buttonHeight float32 = 50

func (t *TTable) newButtonSet() *TButtonSet {
	//fmt.Println("newButtonSet()")
	b := &TButtonSet{}

	b.button1 = widget.NewButton("BET", processButton(t, "BET"))
	b.button1.Importance = widget.HighImportance
	b.button2 = widget.NewButton("MAXBET", processButton(t, "MAXBET"))
	b.button2.Importance = widget.HighImportance
	b.button3 = widget.NewButton("RETRY", processButton(t, "RETRY"))
	b.button3.Importance = widget.HighImportance
	b.button4 = widget.NewButton("CREDIT", processButton(t, "CREDIT"))
	b.button4.Importance = widget.HighImportance
	b.button5 = widget.NewButton("DOUBLE", processButton(t, "DOUBLE"))
	b.button5.Importance = widget.HighImportance
	b.button6 = widget.NewButton("RED", processButton(t, "RED"))
	b.button6.Importance = widget.HighImportance
	b.button7 = widget.NewButton("BLACK", processButton(t, "BLACK"))
	b.button7.Importance = widget.HighImportance

	b.container = container.NewGridWithColumns(7, b.button1, b.button2,
		b.button3, b.button4, b.button5, b.button6, b.button7)

	return b
}

func processButton(t *TTable, s string) func() {
	//fmt.Println("processButton()")
	switch s {
	case "BET":
		return func() {
			if !t.game.Flag1 && !t.game.Flag2 { // State 1
				if t.game.bets < 10 {
					t.game.credits--
					t.game.bets++
					t.game.updatePanelValues(t.game.bets, t.game.credits)
				}
			}
		}
	case "MAXBET":
		return func() {
			if !t.game.Flag1 && !t.game.Flag2 { // State 1
				if t.game.bets < 10 {
					t.game.credits = t.game.credits - (10 - t.game.bets)
					t.game.bets = t.game.bets + (10 - t.game.bets)
					t.game.updatePanelValues(t.game.bets, t.game.credits)
				}
			}
		}
	case "RETRY":
		return func() {
			if t.game.Flag1 && !t.game.Flag2 { // State 2
				t.game.DrawAgain()
				t.game.CheckPrizes()
				t.game.updateWins(t.game.wins)
				t.Refresh()
				if t.game.Flag3 {
				} else {
					t.game.endOfPlay()
				}
			}
		}
	case "CREDIT":
		return func() {
			if (t.game.Flag1 || t.game.Flag2) && t.game.Flag3 { // State 3
				t.game.credits = t.game.credits + t.game.wins
				t.game.bets = 0
				t.game.updatePanelValues(t.game.bets, t.game.credits)
				t.game.Panel.fhelp.Text = msg1
				t.game.endOfPlay()
			}
		}
	case "DOUBLE":
		return func() {
			if (t.game.Flag1 || t.game.Flag2) && t.game.Flag3 { // State 3
				t.game.Flag4 = true
				t.game.Panel.fhelp.Text = msg3
			}
		}
	case "RED":
		return func() {
			if t.game.Flag4 { // State 4
				t.game.Card6 = t.game.drawCard()
				if t.game.Card6 == nil { // No more cards in the deck
					return
				}
				t.Refresh()
				if t.game.Card6.Suit == SuitDiamonds || t.game.Card6.Suit == SuitHearts {
					t.game.wins = 2 * t.game.wins
					t.game.updateWins(t.game.wins)
					t.game.Panel.fhelp.Text = msg5
				} else {
					t.game.endOfPlay()
				}
			}
		}
	case "BLACK":
		return func() {
			if t.game.Flag4 { // State 4
				t.game.Card6 = t.game.drawCard()
				if t.game.Card6 == nil { // No more cards in the deck
					return
				}
				t.Refresh()
				if t.game.Card6.Suit == SuitClubs || t.game.Card6.Suit == SuitSpades {
					t.game.wins = 2 * t.game.wins
					t.game.updateWins(t.game.wins)
					t.game.Panel.fhelp.Text = msg5
				} else {
					t.game.endOfPlay()
				}
			}
		}
	}
	return func() {}
}
