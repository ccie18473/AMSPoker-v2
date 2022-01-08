package main

import (
	//"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

// show creates a new game and loads a table rendered in a new window.
func show(app fyne.App) {
	//fmt.Println("show()")
	game := NewGame()

	w := app.NewWindow("AMSPoker v2.0.2 1996, 2021")
	w.SetPadded(false)
	w.SetContent(NewTable(game))
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(720, 540))

	w.Show()
}

func main() {
	//fmt.Println("main()")
	a := app.New()
	a.SetIcon(ResourceIconPng)
	a.Settings().SetTheme(newGameTheme())

	show(a)
	a.Run()
}
