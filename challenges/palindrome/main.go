package main

import (
	"fmt"
	"os"
)


func main() {

	if len(os.Args) < 1 {
		panic("No string given as parameters")
	}

	value := os.Args[1]
	solution := Solution([]byte(value))
	fmt.Println(string(solution))
}
