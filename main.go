package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ChrisCPoirier/chess/board"
	"github.com/ChrisCPoirier/chess/inputs"
	"github.com/ChrisCPoirier/chess/inputs/openai"
	"github.com/joho/godotenv"
	"github.com/notnil/chess"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("AI chess battle simulator")

	board := board.New()

	myWindow.SetContent(board.Grid)

	player1 := openai.New(`OpenAI Jimmy`, `white`)
	player2 := openai.New(`OpenAI John`, `black`)

	players := []inputs.Player{player1, player2}
	game := chess.NewGame()

	go Loop(myWindow, board, game, players)

	myWindow.ShowAndRun()
}

func Loop(window fyne.Window, board *board.Board, game *chess.Game, players []inputs.Player) {
	i := 0
	for game.Outcome() == chess.NoOutcome {
		time.Sleep(time.Millisecond * 100)
		turn := i % 2

		move, err := players[turn].Ask(game.String())

		if err != nil {
			log.Errorf("%s, try again!", err)
			continue
		}

		if err := game.MoveStr(move); err != nil {
			log.Errorf("%s, try again!", err)
			continue
		}

		log.Info(game.String())

		board.LoadFromFEN(game.FEN())
		board.Grid.Refresh()
		i++
	}
}
