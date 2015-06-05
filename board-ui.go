package main

import (
	"github.com/nsf/termbox-go"
	"strconv"
	"fmt"
)

type Board struct{

}

var player1 = termbox.ColorYellow;
var player2 = termbox.ColorGreen;
var neutrino = termbox.ColorWhite;
var bg = termbox.ColorBlack;
var fg = termbox.ColorWhite;


func (b *Board) Move(oldX, oldY, newX, newY int) error{

	if oldX < 1 || oldX > 5 || oldY < 1 || oldY > 5{
		return fmt.Errorf("Old coordinates must be between 1 and 5 inclusive, (%d, %d)", oldX, oldY)
	}
	if newX < 1 || newX > 5 || newY < 1 || newY > 5{
		return fmt.Errorf("New coordinates must be between 1 and 5 inclusive, (%d, %d)", newX, newY)
	}

	buffer := termbox.CellBuffer()
	i,_ := termbox.Size()

	oldCell := buffer[i*oldY + oldX]
	oldBG := oldCell.Bg
	oldFG := oldCell.Fg
	oldCh := oldCell.Ch

	termbox.SetCell(oldX, oldY, 'E', bg, bg)
	termbox.SetCell(newX, newY, oldCh, oldFG, oldBG)
	termbox.Flush()

	return nil
}

func (b *Board) InitializeBoard(){
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func (b *Board) CloseBoard(){
	defer termbox.Close()
}

func (b *Board) DrawBoard() {

	for i := 0; i <= 6; i++{
		for j := 0; j <= 6; j++{
			if (i == 0 || i == 6) && j > 0 && j < 6{
				termbox.SetCell(i, j, rune(strconv.Itoa(j)[0]), fg, bg)
			} else if (j == 0 || j == 6) && i > 0 && i < 6{
				r, err := getRuneFromIndex(i)
				if err != nil {
					panic(err)
				}
				termbox.SetCell(i, j, r, fg, bg)
			} else if j == 1 {
				termbox.SetCell(i, j, 'C', player1, player1)
			} else if j == 5 {
				termbox.SetCell(i, j, 'P', player2, player2)
			} else if j == 3 && i == 3 {
				termbox.SetCell(i, j, 'N', neutrino, neutrino);
			} else {
				termbox.SetCell(i, j, 'E', bg, bg);
			}
		}
	}

	termbox.Flush()
}

func getRuneFromIndex(j int) (rune, error) {
	switch j{
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
