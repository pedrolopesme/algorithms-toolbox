package main

import "fmt"

type Cell struct {
	played bool
	mine   bool
}

func (c *Cell) Print() {
	char := "🟩"
	if c.played {
		if c.mine {
			char = "☠️"
		} else {
			char = "😀"
		}
	}

	fmt.Printf("%s", char)
}

type Board struct {
	cells [][]Cell
}

func (b *Board) Print() {
	columnsSize := len(b.cells)
	rowSize := len(b.cells[0])

	// walk throught the entire matrix to print a cells
	for row := 0; row < rowSize; row++ {
		for column := 0; column < columnsSize; column++ {
			b.cells[row][column].Print()
		}
		fmt.Println()
	}
}

func New(columnsSize, rowSize int) *Board {
	board := Board{}
	board.cells = make([][]Cell, rowSize)

	// initializing all cells
	for row := 0; row < rowSize; row++ {
		board.cells[row] = make([]Cell, columnsSize)
		for column := 0; column < columnsSize; column++ {
			board.cells[row][column] = Cell{}
		}
	}
	return &board
}

func main() {
	board := New(5, 5)
	board.Print()
}
