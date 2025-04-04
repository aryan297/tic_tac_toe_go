package game

import (
	"fmt"

	"github.com/aryan297/tic-tac-toe/game/utils"
)

type Player interface {
	GetMove(board *Board) (int, int)
	GetSymbol() Symbol
	GetName() string
}

type HumanPlayer struct {
	Name   string
	Symbol Symbol
}

func (hp *HumanPlayer) GetMove(board *Board) (int, int) {
	fmt.Printf("%s's Turn (%s). Enter row and column: ", hp.Name, hp.Symbol)
	row := utils.ReadInt("Row")
	col := utils.ReadInt("Col")
	return row, col
}

func (hp *HumanPlayer) GetSymbol() Symbol {
	return hp.Symbol
}

func (hp *HumanPlayer) GetName() string {
	return hp.Name
}
