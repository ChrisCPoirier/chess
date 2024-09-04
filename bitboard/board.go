package bitboard

import (
	_ "image/png" // Register PNG format

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Board struct {
	Grid  *tview.Grid
	cells [][]*tview.TextView
}

func New(rows, cols int) *Board {
	b := &Board{}
	newCell := func() *tview.TextView {

		textView := tview.
			NewTextView().
			SetText(``).
			SetTextAlign(tview.AlignCenter)

		return textView
	}

	b.Grid = tview.NewGrid().
		// SetBorders(true).
		SetSize(rows, cols, 1, 1).
		SetGap(0, 0)

	// b.Grid.SetBorderPadding(0, 0, 0, 0)

	for r := 0; r < rows; r++ {
		row := []*tview.TextView{}
		for c := 0; c < cols; c++ {
			cell := newCell()

			if (r+c)%2 == 0 {
				cell.SetBackgroundColor(tcell.ColorWhite)
			} else {
				cell.SetBackgroundColor(tcell.ColorBlue)
			}
			cell.SetBorderPadding(0, 0, 0, 0)

			row = append(row, cell)
			b.Grid.AddItem(cell, r, c, 1, 1, 0, 0, false)
		}
		b.cells = append(b.cells, row)
	}

	return b
}

func (g *Board) Paint(values [][]int) *Board {
	for r, row := range values {
		for c, _ := range row {
			// g.cells[r][c].SetText(chess.Display[val])
			g.cells[r][c].SetTextStyle(tcell.StyleDefault.Attributes())
		}
	}
	return g
}
