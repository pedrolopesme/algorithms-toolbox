package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
// source: https://www.hackerrank.com/challenges/ctci-ransom-note/problem
func checkMagazine(magazine []string, note []string) bool {
	magazineMap := map[string]int{}
	for i := 0; i < len(magazine); i++ {
		magazineMap[magazine[i]]++
	}

	matchedWords := int(0)
	for i := 0; i < len(note); i++ {
		if counter, contains := magazineMap[note[i]]; contains {
			if counter > 0 {
				magazineMap[note[i]]--
				matchedWords++
			}
		}
	}

	if matchedWords == len(note) {
		fmt.Printf("Yes")
		return true
	} else {
		fmt.Printf("No")
		return false
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
