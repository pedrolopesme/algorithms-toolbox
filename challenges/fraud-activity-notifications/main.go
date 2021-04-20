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
func activityNotifications(expenditure []int32, d int32) int32 {

	// lets use count sort to avoid heavy sorting each
	// iteration on expenditure. First we need to
	// prepare the beginning of d elements on expenditure
	expendCountSort := make([]int32, 200) // by definition, max expenditure is 200
	for i := int32(0); i < d; i++ {
		expendCountSort[expenditure[i]]++
	}

	notifications := int32(0)
	for i := d; i < int32(len(expenditure)); i++ {
		currentExpense := expenditure[i]
		median := getMedian(expendCountSort, d)

		if float32(currentExpense) >= 2.0*median {
			notifications++
		}

		// update count sort
		expendCountSort[currentExpense]++
		expendCountSort[expenditure[i-d]]--
	}

	return notifications
}

// this median works adapted to count sort logic
func getMedian(expenditure []int32, d int32) float32 {
	if d%2 == 0 {
		var middle1, middle2 *int
		count := int32(0)

		for i := 0; i < len(expenditure); i++ {
			count += expenditure[i]
			if middle1 == nil && count >= d/2 {
				middle1 = &i
			}

			if middle2 == nil && count >= d/2+1 {
				middle2 = &i
				break
			}

		}
		return float32((*middle1 + *middle2) / 2.0)
	} else {
		count := int32(0)
		for i := 0; i < len(expenditure); i++ {
			count += expenditure[i]
			if count > d/2 {
				return float32(i)
			}
		}
	}
	return 0.0
}

func activityNotificationsSlow(expenditure []int32, d int32) int32 {
	notifications := int32(0)
	for i := d; i < int32(len(expenditure)); i++ {
		subset := expenditure[i-d : i]
		notificationThreshold := 2 * medianSlow(subset)
		dayExpense := float32(expenditure[d])
		if dayExpense >= notificationThreshold {
			notifications++
		}
	}
	return notifications
}

func medianSlow(numbers []int32) float32 {
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
