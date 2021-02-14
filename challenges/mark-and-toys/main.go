package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func merge(a1, a2 []int32) []int32 {
	sorted := []int32{}
	var i1, i2 int

	if i1 < len(a1) && i2 < len(a2) {
		if a1[i1] < a2[i2] {
			sorted = append(sorted, a1[i1])
			i1++
		} else {
			sorted = append(sorted, a2[i2])
			i2++
		}
	}

	sorted = append(sorted, a1[i1:]...)
	sorted = append(sorted, a2[i2:]...)
	return sorted
}

// using mergeSort
func sortPrices(unsortedPrices []int32) []int32 {
	if len(unsortedPrices) <= 1 {
		return unsortedPrices
	}

	middle := len(unsortedPrices) / 2
	return merge(sortPrices(unsortedPrices[:middle]), sortPrices(unsortedPrices[middle:]))
}

// Complete the maximumToys function below.
// source https://www.hackerrank.com/challenges/mark-and-toys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sort
func maximumToys(prices []int32, k int32) int32 {
	// it will be easier to compare prices if they are
	// sorted
	sortedPrices := sortPrices(prices)
	var maxItems int32

	for i := 0; i < len(sortedPrices); i++ {
		// if the current price is bigger than the amount to spend
		// then thereis not to do
		if sortedPrices[i] > k {
			break
		}

		currQttItems := int32(1)
		currSumPrice := sortedPrices[i]

		// Check if we are not in the
		if i < len(sortedPrices)-1 {
			for j := i + 1; j < len(sortedPrices); j++ {
				if sortedPrices[j]+currSumPrice <= k {
					currQttItems++
					currSumPrice += sortedPrices[j]
				}
			}
		}

		if currQttItems > maxItems {
			maxItems = currQttItems
		}
	}

	return maxItems
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	pricesTemp := strings.Split(readLine(reader), " ")

	var prices []int32

	for i := 0; i < int(n); i++ {
		pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
		checkError(err)
		pricesItem := int32(pricesItemTemp)
		prices = append(prices, pricesItem)
	}

	result := maximumToys(prices, k)

	fmt.Fprintf(writer, "%d\n", result)

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
