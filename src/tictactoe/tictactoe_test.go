package tictactoe

import (
	"testing"
)

func TestXStarts(t *testing.T) {
	ttt := NewTicTacToeGame()
	err := ttt.Play(0, 0, X)

	if err != nil {
		t.Errorf("This move was expected to be successfull")
	}
	if ttt.Board[0][0] != X {
		t.Errorf("Failed to insert X at 0, 0")
	}
	if ttt.State != InProgress {
		t.Errorf("Game was expected to have started status")
	}

	if ttt.Turn != PlayerO {
		t.Errorf("O is expected to play after X")
	}
}

func TestPlayOnTakenCell(t *testing.T) {
	ttt := NewTicTacToeGame()
	ttt.Board = Cells{
		{X, Empty, Empty},
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
	}
	err := ttt.Play(0, 0, X)
	if err == nil {
		t.Errorf("Unexpected")
	}
}

func TestPlayerWinsOnRows(t *testing.T) {
	ttt := new(TicTacToeGame)
	ttt.Board = Cells{
		{X,     X, Empty},
		{O,     O, Empty},
		{Empty, Empty, Empty},
	}
	ttt.Turn = PlayerX
	ttt.State = InProgress
	ttt.Play(0, 2, X)

	if ttt.State != PlayerXWon {
		t.Errorf("Unexpected")
	}
}

func TestPlayerWinsOnCols(t *testing.T) {
	ttt := new(TicTacToeGame)
	ttt.Board = Cells{
		{X,     O, Empty},
		{X,     O, Empty},
		{Empty, Empty, Empty},
	}
	ttt.Turn = PlayerO
	ttt.State = InProgress
	ttt.Play(2, 1, O)

	if ttt.State != PlayerOWon {
		t.Errorf("Unexpected")
	}
}

func TestPlayerWinsOnDiagonal(t *testing.T) {
	ttt := new(TicTacToeGame)
	ttt.Board = Cells{
		{X,     O, Empty},
		{Empty, X, Empty},
		{O,     Empty, Empty},
	}
	ttt.Turn = PlayerX
	ttt.State = InProgress
	ttt.Play(2, 2, X)

	if ttt.State != PlayerXWon {
		t.Errorf("Unexpected")
	}
}


func TestGameDraw(t *testing.T) {
	ttt := new(TicTacToeGame)
	ttt.Board = Cells{
		{X,     O,      X},
		{X,     X,      O},
		{O,     Empty,  O},
	}
	ttt.Turn = PlayerX
	ttt.State = InProgress
	ttt.Play(2, 1, X)

	if ttt.State != GameTied {
		t.Errorf("Unexpected")
	}
}
