package game

type GameStatus int

const (
	Ongoing GameStatus = iota
	Draw
	Win
)
