package main

import (
	"github.com/nsf/termbox-go"
	"github.com/morras/go-neutrino"
)

func main() {

	game := neutrino.NewGame()
	moveChannel, stateChannel := neutrino.StartGame(game)
	defer neutrino.EndGame()

	board := NewBoard(moveChannel, stateChannel)
	defer board.CloseBoard()

	termbox.PollEvent()

	neutrino.MakeMove(neutrino.NewMove(2, 2, 2, 1))

	termbox.PollEvent()

	neutrino.MakeMove(neutrino.NewMove(2, 4, 4, 2))

	termbox.PollEvent()
}
