package main

import (
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ChrisCPoirier/chess/board"
	"github.com/ChrisCPoirier/chess/inputs"
	"github.com/ChrisCPoirier/chess/inputs/anthropic"
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

	hbox := container.New(layout.NewVBoxLayout(), board.Grid, LogPane())

	myWindow.SetContent(hbox)

	player1 := openai.New(`OpenAI`, `white`)
	player2 := anthropic.New(`Anthropic`, `black`)

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

		player := players[turn]

		err := recursiveAsk(player, game, []string{}, 0)

		if err != nil {
			log.WithField(`player`, player.Name()).Error(err)
			continue
		}

		log.WithField(`player`, player.Name()).Info(game.String())

		board.LoadFromFEN(game.FEN())
		board.Grid.Refresh()
		i++
	}

	dialog.ShowConfirm(`game complete!`, game.Outcome().String(), func(bool) { window.Close() }, window)
}

func recursiveAsk(p inputs.Player, game *chess.Game, invalidAttempts []string, depth int) error {
	if depth > 10 {
		return errors.New(`max depth exceeded`)
	}

	move, err := p.Ask(game.String(), invalidAttempts)

	if err != nil {
		log.WithField(`player`, p.Name()).Error(err)
		return recursiveAsk(p, game, invalidAttempts, depth+1)
	}

	if err := game.MoveStr(move); err != nil {
		log.WithField(`player`, p.Name()).Error(err)
		return recursiveAsk(p, game, append(invalidAttempts, move), depth+1)
	}

	return nil
}

func LogPane() *fyne.Container {
	logs := []fyne.CanvasObject{
		widget.NewLabel(`logs...`),
		widget.NewLabel(``),
		widget.NewLabel(``),
		widget.NewLabel(``),
		widget.NewLabel(``),
	}
	logsContaner := container.NewVBox(logs...)
	logsContaner.Resize(fyne.NewSize(400, 400))

	return logsContaner
}
