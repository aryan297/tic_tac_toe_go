package main

import (
	"github.com/aryan297/tic-tac-toe/game/game"
	"github.com/aryan297/tic-tac-toe/game/utils"
)

func main() {
	boardSize := utils.ReadInt("Enter board size")
	winLength := utils.ReadInt("Enter win length")
	playerCount := utils.ReadInt("Enter number of players")

	g := game.NewGame(boardSize, winLength, playerCount)
	g.Start()
}
