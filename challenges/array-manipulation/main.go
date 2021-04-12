package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
// source: https://www.hackerrank.com/challenges/crush/problem
func arrayManipulation2(n int32, queries [][]int32) int64 {
	arr := make([]int64, n)
	maxElement := int64(0)
	for i := 0; i < len(queries); i++ {
		for j := queries[i][0] - 1; j < queries[i][1]; j++ {
			arr[j] += int64(queries[i][2])
			if maxElement < arr[j] {
				maxElement = arr[j]
			}
		}
	}
	return maxElement
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+2)
	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]
		k := queries[i][2]
		arr[a] += int64(k)
		arr[b+1] -= int64(k)
	}

	max := int64(0)
	value := int64(0)
	for i := int32(0); i < n; i++ {
		value += arr[i]
		max = maxVal(max, value)
	}
	return max
}

func maxVal(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
