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

// Complete the activityNotifications function below.
// source: https://www.hackerrank.com/challenges/fraudulent-activity-notifications/problem
// TODO it is slow. Make it faster
func activityNotifications(expenditure []int32, d int32) int32 {
	notifications := int32(0)
	for i := d; i < int32(len(expenditure)); i++ {
		subset := expenditure[i-d : i]
		notificationThreshold := 2 * median(subset)
		dayExpense := float32(expenditure[d])
		if dayExpense >= notificationThreshold {
			notifications++
		}
	}
	return notifications
}

func median(numbers []int32) float32 {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	index := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return float32(numbers[index-1]+numbers[index]) / float32(2)
	}

	return float32(numbers[index])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

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
