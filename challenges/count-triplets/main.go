package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countTriplets function below.
// source: https://www.hackerrank.com/challenges/count-triplets-1/problem
func countTriplets(arr []int64, r int64) int64 {

	// first turn array into a hashmap
	// it will spare us from making heavy searches multiple
	// times in the future to find all matching pairs
	arrMap := make(map[int64]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		arrMap[arr[i]]++
	}

	tripletsTotal := int64(0)
	for i := 0; i < len(arr); i++ {
		firstElement := arr[i]

		// trying to find the second element
		secondElement := firstElement * r
		if secondElementAmount, contains := arrMap[secondElement]; contains {
			thirdElement := secondElement * r

			// as the second element can exist multiple times,
			// we are going to find the third one each time the second appears
			thirdElementAmount := arrMap[thirdElement]
			triplets := secondElementAmount * thirdElementAmount
			tripletsTotal += triplets
		}
	}

	return tripletsTotal
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(nr[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	r, err := strconv.ParseInt(nr[1], 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < int(n); i++ {
		arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItem)
	}

	ans := countTriplets(arr, r)

	fmt.Fprintf(writer, "%d\n", ans)

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
