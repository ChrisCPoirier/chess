package board

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	"fyne.io/fyne/v2"
	"github.com/sirupsen/logrus"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var (
	COLOR_DARK_SQUARE  = color.RGBA{112, 102, 119, 0xFF}
	COLOR_LIGHT_SQUARE = color.RGBA{204, 183, 174, 0xFF}
)

type Board struct {
	Grid *fyne.Container
}

func New() *Board {

	squares := []fyne.CanvasObject{}

	for r := range 8 {
		for c := range 8 {
			color := getSquareColor(r + c)
			squares = append(squares, newSquare(color, ``))
		}
	}

	b := Board{Grid: container.New(layout.NewGridLayout(8), squares...)}
	b.LoadFromFEN(STARTING_POS_FEN)

	return &b
}

const STARTING_POS_FEN = `rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR`

// TODO: This is terrible but simplifies the loop
func pad(s string) string {

	//replace any number with that number of spaces
	// range of 9 will assign i 0 - 8
	for i := range 9 {
		s = strings.ReplaceAll(s, fmt.Sprintf(`%d`, i), strings.Repeat(` `, i))
	}

	//replace any `/` with no string
	s = strings.ReplaceAll(s, `/`, ``)
	return s
}

func (b *Board) LoadFromFEN(fen string) {
	logrus.Infof("parsing FEN: %s", fen)
	//TODO: add validator on `/` count
	//TODO: add validator on piece count
	//TODO: add validator on line count [characters + spaces] <= 8

	//TODO: consider removing this and making this more elegant/efficient
	fen = pad(fen)

	j := 0
	for i := range b.Grid.Objects {
		switch fen[j] {
		//TODO: Might be better to just create this mapping in a ... wait for it... a map
		case 'r':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/dark_rook.svg`)
		case 'n':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/dark_knight.svg`)
		case 'b':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/dark_bishop.svg`)
		case 'p':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/dark_pawn.svg`)
		case 'k':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/dark_king.svg`)
		case 'q':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/dark_queen.svg`)
		case 'R':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/light_rook.svg`)
		case 'N':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/light_knight.svg`)
		case 'B':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/light_bishop.svg`)
		case 'P':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/light_pawn.svg`)
		case 'K':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/light_king.svg`)
		case 'Q':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), `assets/light_queen.svg`)
		case ' ':
			b.Grid.Objects[i] = newSquare(getSquareColor(i), ``)
		}
		j++
	}

}

func newSquare(bgColor color.Color, imgLocation string) fyne.CanvasObject {
	sqr := canvas.NewRectangle(bgColor)

	if imgLocation != `` {
		img := canvas.NewImageFromFile(imgLocation)
		img.SetMinSize(fyne.NewSquareSize(100))
		stk := container.NewStack(sqr, img)

		return stk
	}

	stk := container.NewStack(sqr)
	return stk
}

func getSquareColor(index int) color.RGBA {
	row := math.Floor(float64(index) / 8)
	if (int(row)+index)%2 != 0 {
		return COLOR_DARK_SQUARE
	}
	return COLOR_LIGHT_SQUARE
}
