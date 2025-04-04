package game

type Symbol string

const (
	Empty Symbol = ""
	X     Symbol = "X"
	O     Symbol = "O"
	Z     Symbol = "Z" // Example: extensible
	Smile Symbol = "🙂"
	Star  Symbol = "⭐"
)

func AllSymbols() []Symbol {
	return []Symbol{X, O, Z, Smile, Star}
}
