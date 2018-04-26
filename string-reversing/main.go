package main

import (
	"fmt"
	"os"
)

// Solution for the String Reversing problem. This
// solution only iterates through half of the string to get its reversed
func Solution(val []byte) []byte {

	valLength := len(val)
	halfLength := valLength / 2

	for i := 0;  i < halfLength; i++ {
		temp := (val)[i]

		(val)[i] = (val)[valLength-1-i]
		(val)[valLength-1-i] = temp
	}
	return val
}

func main() {

	if len(os.Args) < 1 {
		panic("No string given as parameters")
	}

	value := os.Args[1]
	solution := Solution([]byte(value))
	fmt.Println(string(solution))
}
