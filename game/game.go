package game

import (
	"fmt"

	"github.com/aryan297/tic-tac-toe/game/utils"
)

type Game struct {
	Board       *Board
	Players     []Player
	CurrentTurn int
	MovesCount  int
	Status      GameStatus
	WinLength   int
}

func NewGame(boardSize int, winLength int, playerCount int) *Game {
	board := NewBoard(boardSize)
	players := make([]Player, playerCount)

	availableSymbols := AllSymbols()
	used := make(map[Symbol]bool)

	for i := 0; i < playerCount; i++ {
		name := fmt.Sprintf("Player %d", i+1)
		symbol := chooseSymbol(name, availableSymbols, used)
		used[symbol] = true
		players[i] = &HumanPlayer{Name: name, Symbol: symbol}
	}

	return &Game{
		Board:       board,
		Players:     players,
		CurrentTurn: 0,
		WinLength:   winLength,
		Status:      Ongoing,
	}
}

func chooseSymbol(playerName string, symbols []Symbol, used map[Symbol]bool) Symbol {
	for {
		fmt.Printf("Available symbols for %s: ", playerName)
		for _, s := range symbols {
			if !used[s] {
				fmt.Printf("%s ", s)
			}
		}
		fmt.Println()
		input := utils.ReadString(fmt.Sprintf("%s, enter your symbol", playerName))
		symbol := Symbol(input)
		if contains(symbols, symbol) && !used[symbol] {
			return symbol
		}
		fmt.Println("Invalid or already used symbol. Try again.")
	}
}

func contains(symbols []Symbol, s Symbol) bool {
	for _, sym := range symbols {
		if sym == s {
			return true
		}
	}
	return false
}

func (g *Game) Start() {
	fmt.Println("\nTic Tac Toe Started!")
	g.Board.Print()

	for g.Status == Ongoing {
		player := g.Players[g.CurrentTurn]
		row, col := player.GetMove(g.Board)
		if g.Board.SetMove(row, col, string(player.GetSymbol())) {
			g.MovesCount++
			g.Board.Print()

			if g.Board.CheckWin(string(player.GetSymbol()), g.WinLength) {
				fmt.Printf("%s wins!\n", player.GetName())
				g.Status = Win
				return
			}

			if g.Board.IsFull() {
				fmt.Println("It's a draw!")
				g.Status = Draw
				return
			}

			g.CurrentTurn = (g.CurrentTurn + 1) % len(g.Players)
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}
