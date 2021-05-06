package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.
// source : https://www.hackerrank.com/challenges/frequency-queries/problem
func freqQuery(queries [][]int32) (ret []int32) {
	elementsMap := map[int32]int32{}      // maps element -> frequecy
	freqMap := map[int32]map[int32]bool{} // maps frequency -> elements

	for q := 0; q < len(queries); q++ {
		query := queries[q]
		operation := queries[q][0]

		switch operation {
		case 1:
			element := queries[q][1]
			mapSetter(freqMap, elementsMap[element], element, false) // flagging old frequency to false
			elementsMap[query[1]]++
			mapSetter(freqMap, elementsMap[element], element, true) // flagging new frequency to true
		case 2:
			element := queries[q][1]
			if elementsMap[query[1]] > 0 { // be sure to not decrement an element that doesnt exists
				mapSetter(freqMap, elementsMap[element], element, false)
				elementsMap[query[1]]--
				mapSetter(freqMap, elementsMap[element], element, true)
			}
		case 3:
			freq := queries[q][1]
			found := false
			for _, exists := range freqMap[freq] {
				if exists {
					ret = append(ret, int32(1))
					found = true
					break
				}
			}
			if !found {
				ret = append(ret, int32(0))
			}
		}
	}
	return
}

func mapSetter(m map[int32]map[int32]bool, key, element int32, val bool) {
	if m[key] == nil {
		m[key] = make(map[int32]bool)
	}
	m[key][element] = val
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
