package tictactoe

import (
	"errors"
	"encoding/json"
)

// A single cell in the board
type Cell uint8

// Possible cell values
const (
	Empty Cell = iota
	X
	O
)

func (c Cell) String() string {
	switch c {
		case Empty:
			return " "
		case X:
			return "X"
		case O:
			return "O"
		default:
			panic("Unknown Cell Type")
	}
}

func (c Cell) MarshalJSON() ([]byte, error) {
	// that's probably not the best thing ever
	return json.Marshal(c.String())
}

func (c *Cell) UnmarshalJSON(b []byte) error {
	var cell_str string
	err := json.Unmarshal(b, &cell_str)
	if err != nil {
		return err
	}
	switch cell_str {
		case "X":
			*c = X
		case "O":
			*c = O
		default:
			return errors.New("Wrong value for cell")
	}
	return nil
}

// Who's next to play
type PlayerTurn uint8

const (
	Any PlayerTurn = iota
	PlayerX
	PlayerO
)

func (pt PlayerTurn) String() string {
	switch pt {
		case Any:
			return "Any"
		case PlayerX:
			return "Player X"
		case PlayerO:
			return "Player O"
		default:
			panic("Unknown player type")
	}
}

func (pt PlayerTurn) MarshalJSON() ([]byte, error) {
	// that's probably not the best thing ever
	return json.Marshal(pt.String())
}

//Game state
type GameState uint8

const (
	InProgress GameState = iota
	PlayerXWon
	PlayerOWon
	GameTied
)

func (s GameState) String() string {
	switch s {
		case InProgress:
			return "In progress"
		case PlayerXWon:
			return "Player X Won"
		case PlayerOWon:
			return "Player O Won"
		default:
			panic("Unknown game state")
	}
}

func (s GameState) MarshalJSON() ([]byte, error) {
	// that's probably not the best thing ever
	return json.Marshal(s.String())
}

// Game Board
type Cells [3][3]Cell


func (b Cells) GetState() GameState {
	xWon := checkAllRows(b, X) || checkAllCols(b, X) || checkAllDiagonals(b, X)
	oWon := checkAllRows(b, O) || checkAllCols(b, O) || checkAllDiagonals(b, O)
	emptyCellsLeft := checkEmptyCellsLeft(b)

	switch {
		case xWon && !oWon:
			return PlayerXWon
		case oWon && !xWon:
			return PlayerOWon
		case !emptyCellsLeft:
			return GameTied
		default:
			return InProgress
	}
}

func checkAllRows(b Cells, c Cell) (bool) {
	return (b[0][0] == c && b[0][1] == c && b[0][2] == c) ||
		(b[1][0] == c && b[1][1] == c && b[1][2] == c) ||
		(b[2][0] == c && b[2][1] == c && b[2][2] == c)
}

func checkAllCols(b Cells, c Cell) (bool) {
	return (b[0][0] == c && b[1][0] == c && b[2][0] == c) ||
		(b[0][1] == c && b[1][1] == c && b[2][1] == c) ||
		(b[0][2] == c && b[1][2] == c && b[2][2] == c)
}

func checkAllDiagonals(b Cells, c Cell) (bool) {
	return (b[0][0] == c && b[1][1] == c && b[2][2] == c) ||
		(b[2][0] == c && b[1][1] == c && b[0][2] == c)
}

func checkEmptyCellsLeft(b Cells) (bool) {
	for _, col := range b {
		for _, cell := range col {
			if cell == Empty {
				return true
			}
		}
	}
	return false
}

func getPlayerByMove(move Cell) PlayerTurn {
	switch move {
		case X:
			return PlayerX
		case O:
			return PlayerO
		default:
			panic("Unknown move type")
	}
}
