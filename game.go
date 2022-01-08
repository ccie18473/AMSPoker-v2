package main

import (
	//"fmt"
	"strconv"
	"time"
)

// Game represents a full AMSPoker game
type TGame struct {
	Deck *TDeck

	Card1, Card2, Card3, Card4, Card5, Card6 *TCard

	Hand *TDeck

	ButtonSet *TButtonSet

	Panel *TPanel

	Flag1, Flag2, Flag3, Flag4 bool // Flag1 - Deck Tapped, Flag2 - Retry Tapped, Flag3 - Winner, Flag4 - Doubling

	bets, wins, credits int
}

// Initial credits

const CREDITS = 100

func (g *TGame) deal() {
	//fmt.Println("deal()")
}

//	Royal Flush
func RoyalFlush(nmbr *[52]int) bool {
	for i := 0; i < 4; i++ {
		if (nmbr[13*i] == 1) &&
			(nmbr[13*i+12] == 1) &&
			(nmbr[13*i+11] == 1) &&
			(nmbr[13*i+10] == 1) &&
			(nmbr[13*i+9] == 1) {
			return true
		}
	}
	return false
}

// StraightFlush
func StraightFlush(nmbr *[52]int) bool {
	for j := 0; j < 4; j++ {
		for k := 0; k < 10; k++ {
			if (nmbr[j*13+k] == 1) &&
				(nmbr[j*13+k+1] == 1) &&
				(nmbr[j*13+k+2] == 1) &&
				(nmbr[j*13+k+3] == 1) &&
				(nmbr[j*13+k+4] == 1) {
				return true
			}
		}
	}
	return false
}

// FourOfAKind
func FourOfAKind(nmbr *[52]int) bool {
	for i := 0; i < 13; i++ {
		if (nmbr[i] == 1) &&
			(nmbr[i+13] == 1) &&
			(nmbr[i+26] == 1) &&
			(nmbr[i+39] == 1) {
			return true
		}
	}
	return false
}

// FullHouse
func FullHouse(nmbr *[52]int) bool {
	var trio bool = false
	var pair bool = false
	var a, b, c int

	for i := 0; i < 13; i++ {
		if (nmbr[i] == 1) &&
			(nmbr[i+13] == 1) &&
			(nmbr[i+26] == 1) {
			trio = true
			nmbr[i] = 0
			nmbr[i+13] = 0
			nmbr[i+26] = 0
			a = i
			b = i + 13
			c = i + 26
		} else if (nmbr[i] == 1) &&
			(nmbr[i+13] == 1) &&
			(nmbr[i+39] == 1) {
			trio = true
			nmbr[i] = 0
			nmbr[i+13] = 0
			nmbr[i+39] = 0
			a = i
			b = i + 13
			c = i + 39
		} else if (nmbr[i] == 1) &&
			(nmbr[i+26] == 1) &&
			(nmbr[i+39] == 1) {
			trio = true
			nmbr[i] = 0
			nmbr[i+26] = 0
			nmbr[i+39] = 0
			a = i
			b = i + 26
			c = i + 39
		} else if (nmbr[i+13] == 1) &&
			(nmbr[i+26] == 1) &&
			(nmbr[i+39] == 1) {
			trio = true
			nmbr[i+13] = 0
			nmbr[i+26] = 0
			nmbr[i+39] = 0
			a = i + 13
			b = i + 26
			c = i + 39
		}
	}

	for j := 0; j < 13; j++ {
		if (nmbr[j] == 1) && (nmbr[j+13] == 1) {
			pair = true
		} else if (nmbr[j] == 1) &&
			(nmbr[j+26] == 1) {
			pair = true
		} else if (nmbr[j] == 1) &&
			(nmbr[j+39] == 1) {
			pair = true
		} else if (nmbr[j+13] == 1) &&
			(nmbr[j+26] == 1) {
			pair = true
		} else if (nmbr[j+13] == 1) &&
			(nmbr[j+39] == 1) {
			pair = true
		} else if (nmbr[j+26] == 1) &&
			(nmbr[j+39] == 1) {
			pair = true
		}
	}
	if trio {
		nmbr[a] = 1
		nmbr[b] = 1
		nmbr[c] = 1
	}
	if trio && pair {
		return true
	} else {
		return false
	}
}

// Flush
func Flush(nmbr *[52]int) bool {
	var count int = 0
	for j := 0; j < 4; j++ {
		for k := 0; k < 13; k++ {
			if nmbr[13*j+k] == 1 {
				count++
			}
		}
		if count == 5 {
			return true
		} else {
			count = 0
		}
	}
	return false
}

// Straight
func Straight(nmbr *[52]int) bool {
	var i int = 0
	var temp int
	var array [5]int

	for j := 0; j < 52; j++ {
		if nmbr[j] == 1 {
			array[i] = j
			i++
		}
	}
	for k := 0; k < 5; k++ {
		array[k] = array[k] % 13
		if array[k] == 0 {
			array[k] = 13
		}
	}
	for a := 0; a < 4; a++ {
		for b := a + 1; b < 5; b++ {
			if array[b] < array[a] {
				temp = array[a]
				array[a] = array[b]
				array[b] = temp
			}
		}
	}
	if (array[4] == array[3]+1) &&
		(array[3] == array[2]+1) &&
		(array[2] == array[1]+1) &&
		(array[1] == array[0]+1) {
		return true
	}
	if (array[4] == 12) &&
		(array[3] == 11) &&
		(array[2] == 10) &&
		(array[1] == 9) &&
		(array[0] == 0) {
		return true
	}
	return false
}

// ThreeOfAKind
func ThreeOfAKind(nmbr *[52]int) bool {
	for i := 0; i < 13; i++ {
		if ((nmbr[i] == 1) &&
			(nmbr[i+13] == 1) &&
			(nmbr[i+26] == 1)) ||
			((nmbr[i] == 1) &&
				(nmbr[i+13] == 1) &&
				(nmbr[i+39] == 1)) ||
			((nmbr[i] == 1) &&
				(nmbr[i+26] == 1) &&
				(nmbr[i+39] == 1)) ||
			((nmbr[i+13] == 1) &&
				(nmbr[i+26] == 1) &&
				(nmbr[i+39] == 1)) {
			return true
		}
	}
	return false
}

// TwoPair
func TwoPair(nmbr *[52]int) bool {
	var count int = 0

	for i := 0; i < 13; i++ {
		if (nmbr[i] == 1) &&
			(nmbr[i+13] == 1) {
			count++
		} else if (nmbr[i] == 1) &&
			(nmbr[i+26] == 1) {
			count++
		} else if (nmbr[i] == 1) &&
			(nmbr[i+39] == 1) {
			count++
		} else if (nmbr[i+13] == 1) &&
			(nmbr[i+26] == 1) {
			count++
		} else if (nmbr[i+13] == 1) &&
			(nmbr[i+39] == 1) {
			count++
		} else if (nmbr[i+26] == 1) &&
			(nmbr[i+39] == 1) {
			count++
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}

// PairOfAces
func PairOfAces(nmbr *[52]int) bool {
	if (nmbr[0] == 1) &&
		(nmbr[13] == 1) {
		return true
	} else if (nmbr[0] == 1) &&
		(nmbr[26] == 1) {
		return true
	} else if (nmbr[0] == 1) &&
		(nmbr[39] == 1) {
		return true
	} else if (nmbr[13] == 1) &&
		(nmbr[26] == 1) {
		return true
	} else if (nmbr[13] == 1) &&
		(nmbr[39] == 1) {
		return true
	} else if (nmbr[26] == 1) &&
		(nmbr[39] == 1) {
		return true
	} else {
		return false
	}
}

func (g *TGame) endOfPlay() {
	//fmt.Println("endOfPlay()")
	g.Panel.fhelp.Text = msg1
	
	g.Flag1 = false // Deck Flag
	g.Flag2 = false // Retry Flag
	g.Flag3 = false // Checkprizes Flag
	g.Flag4 = false // Double Flag
	
	g.bets = 0
	g.updatePanelValues(0, g.credits)
	
	g.wins = 0
	g.updateWins(0)
	
	for i := 0; i < 9; i++ {
		g.Panel.fprize[i].TextStyle.Bold = false
		g.Panel.fprize[i].Refresh()
		g.Panel.fvalue[i].TextStyle.Bold = false
		g.Panel.fvalue[i].Refresh()
	}
}

func (g *TGame) updateWins(wins int) {
	//fmt.Println("updateWins()")
	g.Panel.fwins.Text = strconv.Itoa(wins)
}

func (g *TGame) updatePanelValues(bets, credits int) {
	//fmt.Println("updatePanelValues()")
	g.Panel.fbets.Text = strconv.Itoa(bets)
	g.Panel.fcredits.Text = strconv.Itoa(credits)

	g.Panel.fvalue[0].Text = strconv.Itoa(bets)
	g.Panel.fvalue[1].Text = strconv.Itoa(bets * 2)
	g.Panel.fvalue[2].Text = strconv.Itoa(bets * 3)
	g.Panel.fvalue[3].Text = strconv.Itoa(bets * 5)
	g.Panel.fvalue[4].Text = strconv.Itoa(bets * 8)
	g.Panel.fvalue[5].Text = strconv.Itoa(bets * 10)
	g.Panel.fvalue[6].Text = strconv.Itoa(bets * 25)
	g.Panel.fvalue[7].Text = strconv.Itoa(bets * 80)
	g.Panel.fvalue[8].Text = strconv.Itoa(bets * 500)
}

func (g *TGame) CheckPrizes() bool {
	//fmt.Println("CheckPrizes()")
	var nmbr [52]int = [52]int{}
	for i, _ := range g.Hand.Cards {
		nmbr[g.Hand.Cards[i].Value+int(13*g.Hand.Cards[i].Suit)] = 1
	}

	if RoyalFlush(&nmbr) {
		g.wins = g.bets * 500
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[8].TextStyle.Bold = true
		g.Panel.fprize[8].Refresh()
		g.Panel.fvalue[8].TextStyle.Bold = true
		g.Panel.fvalue[8].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if StraightFlush(&nmbr) {
		g.wins = g.bets * 80
		g.Flag3 = true
		g.Panel.fprize[7].TextStyle.Bold = true
		g.Panel.fprize[7].Refresh()
		g.Panel.fvalue[7].TextStyle.Bold = true
		g.Panel.fvalue[7].Refresh()
		g.Panel.fprize[7].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if FourOfAKind(&nmbr) {
		g.wins = g.bets * 25
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[6].TextStyle.Bold = true
		g.Panel.fprize[6].Refresh()
		g.Panel.fvalue[6].TextStyle.Bold = true
		g.Panel.fvalue[6].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if FullHouse(&nmbr) {
		g.wins = g.bets * 10
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[5].TextStyle.Bold = true
		g.Panel.fprize[5].Refresh()
		g.Panel.fvalue[5].TextStyle.Bold = true
		g.Panel.fvalue[5].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if Flush(&nmbr) {
		g.wins = g.bets * 8
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[4].TextStyle.Bold = true
		g.Panel.fprize[4].Refresh()
		g.Panel.fvalue[4].TextStyle.Bold = true
		g.Panel.fvalue[4].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if Straight(&nmbr) {
		g.wins = g.bets * 5
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[3].TextStyle.Bold = true
		g.Panel.fprize[3].Refresh()
		g.Panel.fvalue[3].TextStyle.Bold = true
		g.Panel.fvalue[3].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if ThreeOfAKind(&nmbr) {
		g.wins = g.bets * 3
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[2].TextStyle.Bold = true
		g.Panel.fprize[2].Refresh()
		g.Panel.fvalue[2].TextStyle.Bold = true
		g.Panel.fvalue[2].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if TwoPair(&nmbr) {
		g.wins = g.bets * 2
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[1].TextStyle.Bold = true
		g.Panel.fprize[1].Refresh()
		g.Panel.fvalue[1].TextStyle.Bold = true
		g.Panel.fvalue[1].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else if PairOfAces(&nmbr) {
		g.wins = g.bets
		g.Flag3 = true
		g.updateWins(g.wins)
		g.Panel.fprize[0].TextStyle.Bold = true
		g.Panel.fprize[0].Refresh()
		g.Panel.fvalue[0].TextStyle.Bold = true
		g.Panel.fvalue[0].Refresh()
		if g.Flag1 && !g.Flag2 {
			g.Panel.fhelp.Text = msg4
		} else if g.Flag2 {
			g.Panel.fhelp.Text = msg5
		}
	} else {
		g.Flag3 = false
	}
	return g.Flag3
}

// ResetDraw resets the draw pile to be completely available (no cards drawn)
func (g *TGame) ResetDraw() {
	//fmt.Println("ResetDraw()")
	for ; len(g.Deck.Cards) > 0; g.DrawFive() {
	}

	// Reset the draw pile
	g.DrawFive()
}

// Remove card from Hand to keep 5 total
func (g *TGame) removeCard(c *TCard) {
	//fmt.Println("removeCard()")
	for i, h := range g.Hand.Cards {
		if (c.Value + int(c.Suit*13)) == (h.Value + int(h.Suit*13)) {
			g.Hand.Cards = append(g.Hand.Cards[:i], g.Hand.Cards[i+1:]...)
		}
	}
}

func (g *TGame) drawCard() *TCard {
	//fmt.Println("drawCard()")
	if len(g.Deck.Cards) == 0 {
		g.Panel.fhelp.Text = msg6 // No more cards in the Deck
		return nil
	}

	// Pop from Deck
	popped := g.Deck.Pop()
	popped.FaceUp = true
	// Push to Hand
	g.Hand.Push(popped)
	return popped
	return nil
}

// Draw again after removing the selected cards
func (g *TGame) DrawAgain() {
	//fmt.Println("DrawAgain")
	if g.Card1.Selected {
		g.removeCard(g.Card1)
		g.Card1 = g.drawCard()
	}
	if g.Card2.Selected {
		g.removeCard(g.Card2)
		g.Card2 = g.drawCard()
	}
	if g.Card3.Selected {
		g.removeCard(g.Card3)
		g.Card3 = g.drawCard()
	}
	if g.Card4.Selected {
		g.removeCard(g.Card4)
		g.Card4 = g.drawCard()
	}
	if g.Card5.Selected {
		g.removeCard(g.Card5)
		g.Card5 = g.drawCard()
	}

	g.Flag2 = true // Retry tapped
}

// DrawFive draws five cards from the deck and adds them to the draw pile(s).
// If there are no cards available to be drawn it will cycle back to the beginning and draw the first three.
func (g *TGame) DrawFive() {
	//fmt.Println("DrawFive()")
	g.Deck = NewShuffledDeck()
	g.Hand = &TDeck{}

	g.Card1 = g.drawCard()
	g.Card2 = g.drawCard()
	g.Card3 = g.drawCard()
	g.Card4 = g.drawCard()
	g.Card5 = g.drawCard()

	g.Card6 = &TCard{99, 3, false, false} // For doubling game

	g.Flag1 = true // Deck tapped

	g.Panel.fhelp.Text = msg2
}

func NewGame() *TGame {
	//fmt.Println("NewGame()")
	return NewGameFromSeed(time.Now().UnixNano())
}

// NewGameFromSeed starts a new AMSPoker game and draws to the standard configuration.
// The randomness of the desk is seeded using the specified value.
func NewGameFromSeed(seed int64) *TGame {
	//fmt.Println("NewGameFromSeed()")
	g := &TGame{}
	g.Deck = NewShuffledDeckFromSeed(seed)
	g.credits = CREDITS
	g.bets = 0

	g.deal()
	return g
}
