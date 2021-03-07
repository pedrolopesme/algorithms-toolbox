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
// source: https://www.hackerrank.com/challenges/new-year-chaos/problem
// returns the number of bribes or -1 if a too chaotic condition
// was found
func minimumBribes2(q []int32) int32 {

	// this map holds the people that were
	// shifted (or accepted a bribe). This is used to detect the reason
	// behind a person is out of its expected position: is that person
	// a briber or just someone that took a bribe?
	shiftedPeople := make(map[int32]bool)

	totalBribes := int32(0)

	maxAllowedBribes := int32(2)

	for i := int32(0); i < int32(len(q)); i++ {
		expectedPerson := (i + 1)
		actualPerson := q[i]

		// try to confirm if we've a problem with current person position
		if expectedPerson != actualPerson {

			// if we know the the current person is out of its position
			// because it had accepted a bribe, we won't count him/her
			// as a briber ...
			if _, ok := shiftedPeople[actualPerson]; ok {
				delete(shiftedPeople, actualPerson)
				continue
			} else {
				// we found a briber! Now, how many positions
				// him/her have bought?
				bribesCounter := abs(actualPerson - expectedPerson)

				// check too chaotic situations
				if bribesCounter > maxAllowedBribes {
					fmt.Println("Too chaotic")
					return int32(-1)
				}

				// counting total bribes
				totalBribes += bribesCounter

				// storing the expected person (we gone meet him/her
				// out of its position in one of further positions)
				shiftedPeople[expectedPerson] = true
				if bribesCounter > 1 {
					// if the briber has bought more than one position
					// we gonna find the person after the next person
					// (I know, I'm just counting with up to two bribings right now)
					shiftedPeople[expectedPerson+1] = true
				}
			}
		}
	}

	fmt.Println(totalBribes)
	return totalBribes
}

func abs(a int32) int32 {
	if a < 0 {
		return a * -1
	}
	return a
}

func minimumBribes(q []int32) int32 {

	totalBribes := int32(0)
	next := int32(1)
	nextNext := int32(2)
	nextNextNext := int32(3)

	for i := int32(0); i < int32(len(q)); i++ {

		if q[i] == next {
			// natural order, just keep moving
			next = nextNext
			nextNext = nextNextNext
			nextNextNext++
		} else if q[i] == nextNext {
			// we found a briber that step up one position
			totalBribes++
			nextNext = nextNextNext
			nextNextNext++
		} else if q[i] == nextNextNext {
			totalBribes += 2
			nextNextNext++
		} else {
			fmt.Println("Too chaotic")
			return int32(-1)
		}
	}
	fmt.Println(totalBribes)
	return totalBribes
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
