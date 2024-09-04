package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ChrisCPoirier/chess/board"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Grid Layout")

	board := board.New()

	myWindow.SetContent(board.Grid)

	go Loop(myWindow, board)
	myWindow.ShowAndRun()
}

func Loop(window fyne.Window, b *board.Board) {
	i := 0
	for {
		//randomize board
		if i%2 == 0 {
			b.LoadFromFEN(`4r1k1/p1p2ppp/2q2b2/4Bb2/8/2N5/PPP2PPP/R3Q1K1`)
		} else {
			b.LoadFromFEN(board.STARTING_POS_FEN)
		}

		b.Grid.Refresh()

		time.Sleep(time.Second * 5)
		i++
	}
}
