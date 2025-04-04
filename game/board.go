package game

import "fmt"

type Board struct {
	Size  int
	Cells [][]*Cell
}

func NewBoard(size int) *Board {
	cells := make([][]*Cell, size)
	for i := 0; i < size; i++ {
		cells[i] = make([]*Cell, size)
		for j := 0; j < size; j++ {
			cells[i][j] = NewCell(i, j)
		}
	}
	return &Board{Size: size, Cells: cells}
}

func (b *Board) Print() {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Cells[i][j].IsEmpty() {
				fmt.Print(" . ")
			} else {
				fmt.Printf(" %s ", b.Cells[i][j].Value)
			}
		}
		fmt.Println()
	}
}

func (b *Board) SetMove(row, col int, symbol string) bool {
	if row >= 0 && row < b.Size && col >= 0 && col < b.Size && b.Cells[row][col].IsEmpty() {
		b.Cells[row][col].Value = symbol
		return true
	}
	return false
}

func (b *Board) CheckWin(symbol string, winLength int) bool {
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}} // right, down, diag, anti-diag

	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			for _, d := range directions {
				count := 0
				for k := 0; k < winLength; k++ {
					x, y := i+d[0]*k, j+d[1]*k
					if x < 0 || x >= b.Size || y < 0 || y >= b.Size || b.Cells[x][y].Value != symbol {
						break
					}
					count++
				}
				if count == winLength {
					return true
				}
			}
		}
	}
	return false
}

func (b *Board) IsFull() bool {
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			if b.Cells[i][j].IsEmpty() {
				return false
			}
		}
	}
	return true
}
