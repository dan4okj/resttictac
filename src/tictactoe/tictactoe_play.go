package tictactoe

import "errors"

type TicTacToeGame struct {
	Board Cells
	Turn PlayerTurn
	State GameState
}

func NewTicTacToeGame() (ttt *TicTacToeGame) {
	ttt = new(TicTacToeGame)
	ttt.Board = Cells{
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
		{Empty, Empty, Empty},
	}
	ttt.Turn = Any
	ttt.State = InProgress
	return
}

func (ttt *TicTacToeGame) Play(col, row int, move Cell) (error) {
	if ttt.State != InProgress {
		return errors.New("Game Finished")
	}
    if ttt.Board[col][row] != Empty {
        return errors.New("Spot taken")
    }
    if move != X && move != O {
        return errors.New("Invalid move")
    }

    player := getPlayerByMove(move)

    if ttt.Turn == Any {
        ttt.Turn = player
    }

    if ttt.Turn != player {
        return errors.New("Not your move")
    }

    ttt.Board[col][row] = move

    ttt.State = ttt.Board.GetState()
	if ttt.State == InProgress{
        if player == PlayerX {
            ttt.Turn = PlayerO
        } else {
            ttt.Turn = PlayerX
        }
    }
    return nil

}
