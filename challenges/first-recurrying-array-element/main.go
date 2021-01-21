package main

import "fmt"

// convention: -1 is repeated element was found,
// otherwise, I'll return the element itself
func FindFirstRecurrying(arr []int) int {
	// validations
	// TODO add a test case for this scenario
	if len(arr) <= 1 {
		return -1
	}

	numberMap := map[int]int{}

	for _, val := range arr {
		if _, ok := numberMap[val]; ok {
			return val
		} else {
			numberMap[val] = 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 2, 3, 4, 5}
	fmt.Println(FindFirstRecurrying(arr))
}
