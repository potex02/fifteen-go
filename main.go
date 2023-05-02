package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/potex02/fifteen-go/game"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//game.Test()
	app := app.New()
	window := app.NewWindow("Minesweeper")
	cells := game.InitializeCells()

	window.SetIcon(theme.ComputerIcon())
	window.Resize(fyne.NewSize(250, 250))
	game.AddElements(window, cells)
	window.ShowAndRun()

}
