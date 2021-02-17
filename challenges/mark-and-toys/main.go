package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func findMinValue(prices []int32, greaterThan int32) (minValue int32) {
	gotFirst := false
	for i := 0; i < len(prices); i++ {
		currValue := prices[i]
		if currValue <= greaterThan {
			continue
		}
		if !gotFirst {
			minValue = currValue
			gotFirst = true
			continue
		}

		if currValue < minValue {
			minValue = currValue
		}
	}
	return
}

// Complete the maximumToys function below.
// source https://www.hackerrank.com/challenges/mark-and-toys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sort
func maximumToys(prices []int32, k int32) (maxToys int32) {
	lastValue := int32(0)
	for k >= 0 {
		lastValue = findMinValue(prices, lastValue)
		k -= lastValue
		if k >= 0 {
			maxToys++
		} else {
			return
		}
	}
	return
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
