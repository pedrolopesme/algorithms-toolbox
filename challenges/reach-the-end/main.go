package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'reachTheEnd' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING_ARRAY grid
 *  2. INTEGER maxTime
 */

func reachTheEnd(grid []string, maxTime int32) string {
	path := findShortestPath(grid)

	// as we are looking for transitions and not elements,
	// lets just ignore first element
	transitions := int32(len(path) - 1)

	if transitions > 0 && transitions <= maxTime {
		return "Yes"
	}

	return "No"
}

// based on BFS, returns the shortest path from the beginning to
// the end of the grid
func findShortestPath(grid []string) (path []int32) {
	current := int32(0)
	queue := []int32{current}
	visited := make([]bool, calculateGridSize(grid))
	visited[current] = true

	for len(queue) > 0 {
		var item int32
		item, queue = queue[0], queue[1:]
		path = append(path, item)

		for _, neighbor := range getNeighbors(item, grid) {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}

	}
	return
}

// returns the total of elements in the grid
func calculateGridSize(grid []string) int32 {
	return int32(len(grid[0]) * len(grid))
}

// returns an array with neighbors
func getNeighbors(current int32, grid []string) (path []int32) {
	row, column := transformElementToPosition(current, grid)
	rowLimit := int32(len(grid)) - 1
	columnLimit := int32(len(grid[0])) - 1

	for x := max(0, row-1); x <= min(row+1, rowLimit); x++ {
		for y := max(0, column-1); y <= min(column+1, columnLimit); y++ {
			// lets ignore diagonals
			// top
			if x == row-1 || x == row+1 {
				if y == column-1 || y == column+1 {
					continue
				}
			}

			if x != row || y != column {
				if cellIsEligible(x, y, grid) {
					path = append(path, transformPositionToElement(x, y, grid))
				}
			}
		}
	}
	return
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// it adapts the string-based grid into an int-based
// matrix by transforming chars in a row into a positional
// array (eg [".##", "###"] turns into [[0,1,2],[3,4,5]].
func transformElementToPosition(element int32, grid []string) (row, column int32) {
	if element == 0 {
		return 0, 0
	}

	elementsPerRow := int32(len(grid[0]))
	row = int32(math.Floor(float64(element) / float64(elementsPerRow)))
	column = element % elementsPerRow
	return
}

func transformPositionToElement(row, column int32, grid []string) int32 {
	elementsPerRow := int32(len(grid[0]))
	firstElement := row * elementsPerRow
	return firstElement + column
}

func cellIsEligible(row, column int32, grid []string) bool {
	return string(grid[row][column]) == "."
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gridCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var grid []string

	for i := 0; i < int(gridCount); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	maxTimeTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	maxTime := int32(maxTimeTemp)

	result := reachTheEnd(grid, maxTime)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
