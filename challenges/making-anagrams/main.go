package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func stringMapIncrement(m *map[rune]int32, s string) {
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		(*m)[r]++
	}
}

func stringMapDecrement(m *map[rune]int32, s string) {
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		(*m)[r]--
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// Complete the makeAnagram function below.
func makeAnagram(a string, b string) int32 {
	stringsMap := make(map[rune]int32)
	stringMapIncrement(&stringsMap, a) // first I create a map counting strings up
	stringMapDecrement(&stringsMap, b) // then a decrement the map counter with the second string
	// Now stringsMap countains the diff
	var diffChars int32
	for _, v := range stringsMap {
		diffChars += abs(v)
	}

	return diffChars
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	a := readLine(reader)

	b := readLine(reader)

	res := makeAnagram(a, b)

	fmt.Fprintf(writer, "%d\n", res)

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
