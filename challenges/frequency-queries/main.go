package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func InsertSorted(data []int32, e int32) []int32 {
	index := sort.Search(len(data), func(i int) bool { return data[i] > e })
	data = append(data, e)
	copy(data[index+1:], data[index:])
	data[index] = e
	return data
}

func RemoveSorted(data []int32, e int32) []int32 {
	if len(data) == 0 {
		return data
	}

	index := sort.Search(len(data), func(i int) bool { return data[i] > e })
	if index < len(data) {
		copy(data[index+1:], data[index:])
	}

	return data
}

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) (ret []int32) {
	opMap := map[int32]int32{}
	freq := make([]int32, len(queries))

	for q := 0; q < len(queries); q++ {
		query := queries[q]

		switch query[0] {
		case 1:
			freq = RemoveSorted(freq, query[1])
			opMap[query[1]]++
			freq = InsertSorted(freq, query[1])
		case 2:
			if val := opMap[query[1]]; val > 0 {
				freq = RemoveSorted(freq, query[1])
				opMap[query[1]]--
				freq = InsertSorted(freq, query[1])
			}
		case 3:
			i := sort.Search(len(freq), func(i int) bool { return freq[i] >= query[1] })
			if i < len(freq) && freq[i] == query[1] {
				ret = append(ret, int32(1))
			} else {
				ret = append(ret, int32(0))
			}
		}
	}
	return
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
