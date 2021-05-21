package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) []int32 {
	costs := map[int32][]int32{} // map of icecream cost -> array of positions

	for i := int32(0); i < int32(len(cost)); i++ {
		if positions, contains := costs[cost[i]]; contains {
			positions := append(positions, i+int32(1))
			costs[cost[i]] = positions
		} else {
			costs[cost[i]] = []int32{i + int32(1)}
		}
	}

	ids := []int32{}
	for i := 0; i < len(cost); i++ {
		currentCost := cost[i]
		leftMoney := money - currentCost

		// special case
		if currentCost == leftMoney {
			positions := costs[currentCost]
			if len(positions) >= 2 {
				ids = positions[0:2]
				break
			}
		} else if positions, contains := costs[leftMoney]; contains {
			ids = append(ids, costs[currentCost][0])
			ids = append(ids, positions[0])
			break
		}
	}

	if ids[0] > ids[1] {
		temp := ids[0]
		ids[0] = ids[1]
		ids[1] = temp
	}

	for i := 0; i < len(ids); i++ {
		fmt.Printf("%d ", ids[i])
	}
	fmt.Println("")
	return ids
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		moneyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		money := int32(moneyTemp)

		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		costTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cost []int32

		for i := 0; i < int(n); i++ {
			costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
			checkError(err)
			costItem := int32(costItemTemp)
			cost = append(cost, costItem)
		}

		whatFlavors(cost, money)
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
