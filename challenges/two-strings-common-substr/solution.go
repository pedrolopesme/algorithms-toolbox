package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the twoStrings function below.
func createCharsArray(str string) map[byte]bool {
	chars := make(map[byte]bool, 26)
	for i := 0; i < len(str); i++ {
		chars[byte(str[i])] = true
	}
	return chars
}

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	var charsMap map[byte]bool
	var compareWith *string

	if len(s1) < len(s2) {
		charsMap = createCharsArray(s1)
		compareWith = &s2
	} else {
		charsMap = createCharsArray(s2)
		compareWith = &s1
	}
	for i := 0; i < len(*compareWith); i++ {
		if charsMap[byte((*compareWith)[i])] {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
	}

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