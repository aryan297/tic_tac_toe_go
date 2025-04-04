package game

type Cell struct {
	Row   int
	Col   int
	Value string // "X", "O", or ""
}

func NewCell(row, col int) *Cell {
	return &Cell{Row: row, Col: col, Value: ""}
}

func (c *Cell) IsEmpty() bool {
	return c.Value == ""
}
