package main

import (
	"github.com/morras/go-neutrino"
	"github.com/nsf/termbox-go"
)

func main() {

	game := neutrino.NewStandardGame()
	moveChannel, stateChannel := neutrino.StartGame(game)
	defer neutrino.EndGame()

	board := NewBoard(moveChannel, stateChannel)//Should be able to take a game perhaps
	defer board.CloseBoard()

	termbox.PollEvent()

	neutrino.MakeMove(neutrino.NewMove(2, 2, 2, 1))

	termbox.PollEvent()

	neutrino.MakeMove(neutrino.NewMove(2, 4, 4, 2))

	termbox.PollEvent()
}
