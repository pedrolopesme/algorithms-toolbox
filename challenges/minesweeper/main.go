package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func (b *Board) PrintRowHeader(row int) {
	fmt.Printf("%s ", string(rune(row+65)))
}

func (b *Board) PrintColumnHeader() {
	for column := 0; column < len(b.cells[0]); column++ {
		fmt.Printf("%d ", column)
	}
	fmt.Println()
}

func (b *Board) Print() {
	fmt.Print("\033[H\033[2J")
	rowSize := len(b.cells)
	columnsSize := len(b.cells[0])

	// walk through the entire matrix to print cells
	for row := 0; row < rowSize; row++ {
		for column := 0; column < columnsSize; column++ {
			if row == 0 && column == 0 {
				fmt.Print("   ") // <- creating an empty corner at top-left on the board
				b.PrintColumnHeader()
			}

			if column == 0 {
				b.PrintRowHeader(row)
			}
			b.cells[row][column].Print()
		}
		fmt.Println()
	}
}

func New(columnsSize, rowSize int) *Board {
	board := Board{}
	board.cells = make([][]Cell, rowSize)

	// initializing all cells
	// TODO is it really necessary?
	for row := 0; row < rowSize; row++ {
		board.cells[row] = make([]Cell, columnsSize)
		for column := 0; column < columnsSize; column++ {
			board.cells[row][column] = Cell{}
		}
	}
	return &board
}

func main() {
	board := New(10, 5)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		board.Print()
		fmt.Print("-> Type row and column number (e.g. A1) - or exit to quit: \n")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

	}
}
