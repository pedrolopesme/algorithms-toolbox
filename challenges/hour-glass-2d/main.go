package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Given an X/Y position, it'll sum all elements around the center.
func calculateHourglass(arr [][]int32, centerX int, centerY int) int32 {
	var result int32
	result += arr[centerX][centerY]
	result += arr[centerX-1][centerY-1]
	result += arr[centerX-1][centerY]
	result += arr[centerX-1][centerY+1]
	result += arr[centerX+1][centerY-1]
	result += arr[centerX+1][centerY]
	result += arr[centerX+1][centerY+1]
	return result
}

// Complete the hourglassSum function below.
// source: https://www.hackerrank.com/challenges/2d-array/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
func hourglassSum(arr [][]int32) int32 {
	var maxResult int32
	first := true
	for x := 1; x < len(arr)-1; x++ {
		for y := 1; y < len(arr)-1; y++ {
			currResult := calculateHourglass(arr, x, y)
			if first {
				maxResult = currResult
				first = false
			} else if currResult > maxResult {
				maxResult = currResult
			}
		}
	}
	return maxResult
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)
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
