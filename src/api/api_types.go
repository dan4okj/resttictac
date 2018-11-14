package api

import (
	"errors"
	"tictactoe"
)

type PlayReq struct {
	Col int `json:"col"`
	Row int `json:"row"`
	Move tictactoe.Cell `json:"move"`
}

func (pr *PlayReq) Validate() error {
	if pr.Col < 0 || pr.Col > 2 {
		return errors.New("Wrong column value")
	}
	if pr.Row < 0 || pr.Row > 2 {
		return errors.New("Wrong row value")
	}
	return nil
}
