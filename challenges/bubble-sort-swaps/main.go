package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countSwaps function below.
// challenge: https://www.hackerrank.com/challenges/ctci-bubble-sort/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting
func countSwaps(arr []int32) (swaps, first, last int32) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr); i++ {
			if i+1 < len(arr) { // check if we've reached the end of the array
				if arr[i] > arr[i+1] {
					swaps++
					temp := arr[i]
					arr[i] = arr[i+1]
					arr[i+1] = temp
					sorted = false
				}
			}
		}
	}
	first = arr[0]
	last = arr[len(arr)-1]
	fmt.Printf("Array is sorted in %d swaps.\nFirst Element: %d\nLast Element: %d", swaps, first, last)
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	countSwaps(a)
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
