package main

import (
	 "github.com/nsf/termbox-go"
)


func main() {

	board := &Board{}

	board.InitializeBoard()
	defer board.CloseBoard()

	board.DrawBoard()

	termbox.PollEvent()

	board.Move(3, 3, 3, 2)

	termbox.PollEvent()

	board.Move(3, 5, 5, 3)

	termbox.PollEvent()
}

