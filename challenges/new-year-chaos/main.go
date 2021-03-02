package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
// https://www.hackerrank.com/challenges/new-year-chaos/problem
func minimumBribes(q []int32) (result int32) {
	// map containing the briber number and its
	// original position
	bribers := make(map[int32]int)

	for i := len(q) - 1; i >= 0; i-- {
		expectedBriber := int32(i + 1)
		if q[i] != expectedBriber {
			// if the briber was identified
			if originalPosition, ok := bribers[q[i]]; ok {
				// try to identify chaotic briber
				if originalPosition-i >= 3 {
					fmt.Print("Too chaotic")
					result = int32(-1)
					break
				} else {
					result += int32(originalPosition - i)
				}
			} else {
				bribers[q[i]] = i
			}
		}
	}

	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
