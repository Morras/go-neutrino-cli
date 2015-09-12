package main

import (
	"fmt"
	"github.com/morras/go-neutrino"
	"github.com/nsf/termbox-go"
	"strconv"
)

type Board struct {
	moveChannel  <-chan neutrino.Move
	stateChannel <-chan neutrino.State
}

func (self *Board) listenForMoves() {
	for move := range self.moveChannel {
		self.move(move)
	}
}

func (self *Board) listenForStateChanges() {
	for state := range self.stateChannel {
		switch state {
		case neutrino.Player1NeutrinoMove:
			writeMessage("Player ones turn to move neutrino")
			break
		case neutrino.Player1Move:
			writeMessage("Player ones turn to move one of their own pieces")
			break
		case neutrino.Player2NeutrinoMove:
			writeMessage("Player twos turn to move neutrino")
			break
		case neutrino.Player2Move:
			writeMessage("Player twos turn to move one of their own pieces")
			break
		case neutrino.Player1Win:
			writeMessage("Player one wins!")
			break
		case neutrino.Player2Win:
			writeMessage("Player two wins!")
			break
		default:
			writeMessage("Invalid state")
		}
	}
}

var lastMessageLength = 0

func writeMessage(message string) {
	for i := 0; i < lastMessageLength; i++ {
		termbox.SetCell(1+i, 8, ' ', 0, 0)
	}
	for i, r := range message {
		termbox.SetCell(1+i, 8, r, 0, 0)
	}
	lastMessageLength = len(message)
	termbox.Flush()
}

const (
	player1Square  = termbox.ColorYellow
	player2Square  = termbox.ColorGreen
	blankSquare    = ' '
	neutrinoSquare = termbox.ColorWhite
	bg             = termbox.ColorBlack
	fg             = termbox.ColorWhite
)

func NewBoard(moveCh <-chan neutrino.Move, stateCh <-chan neutrino.State) *Board {
	board := &Board{
		moveChannel:  moveCh,
		stateChannel: stateCh,
	}
	board.initializeBoard()
	board.drawBoard()
	go board.listenForMoves()
	go board.listenForStateChanges()
	return board
}

func (self *Board) move(move neutrino.Move) error {
	//the board that is drawn is 1-indexed instead of 0-indexed
	move.FromX++
	move.ToX++
	move.FromY++
	move.ToY++

	if move.FromX < 1 || move.FromX > 5 || move.FromY < 1 || move.FromY > 5 {
		return fmt.Errorf("Old coordinates must be between 1 and 5 inclusive, (%d, %d)", move.FromX, move.FromY)
	}
	if move.ToX < 1 || move.ToX > 5 || move.ToY < 1 || move.ToY > 5 {
		return fmt.Errorf("New coordinates must be between 1 and 5 inclusive, (%d, %d)", move.ToX, move.ToY)
	}

	buffer := termbox.CellBuffer()
	bufferWidth, _ := termbox.Size()

	oldCell := buffer[bufferWidth*int(move.FromY)+int(move.FromX)]
	oldCh := oldCell.Ch
	oldFG := oldCell.Fg
	oldBG := oldCell.Bg

	termbox.SetCell(int(move.ToX), int(move.ToY), oldCh, oldFG, oldBG)
	termbox.SetCell(int(move.FromX), int(move.FromY), ' ', fg, bg)
	termbox.Flush()

	return nil
}

func (self *Board) initializeBoard() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (self *Board) CloseBoard() {
	defer termbox.Close()
}

func (self *Board) drawBoard() {

	for i := 0; i <= 6; i++ {
		for j := 0; j <= 6; j++ {
			if (i == 0 || i == 6) && j > 0 && j < 6 {
				termbox.SetCell(i, j, rune(strconv.Itoa(j)[0]), fg, bg)
			} else if (j == 0 || j == 6) && i > 0 && i < 6 {
				r, err := getRuneFromIndex(i)
				if err != nil {
					panic(err)
				}
				termbox.SetCell(i, j, r, fg, bg)
			} else if j == 1 {
				termbox.SetCell(i, j, ' ', termbox.ColorWhite, player1Square)
			} else if j == 5 {
				termbox.SetCell(i, j, ' ', termbox.ColorWhite, player2Square)
			} else if j == 3 && i == 3 {
				termbox.SetCell(i, j, ' ', termbox.ColorWhite, neutrinoSquare)
			} else {
				termbox.SetCell(i, j, ' ', bg, bg)
			}
		}
	}

	termbox.Flush()
}

func getRuneFromIndex(j int) (rune, error) {
	switch j {
	case 1:
		return 'A', nil
	case 2:
		return 'B', nil
	case 3:
		return 'C', nil
	case 4:
		return 'D', nil
	case 5:
		return 'E', nil
	}
	return 'Q', fmt.Errorf("input is not between 1 and 5,%d", j)
}
