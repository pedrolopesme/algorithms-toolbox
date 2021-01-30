package main

import "fmt"

func bubbleSort(arr []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < len(arr); i++ {
			if i+1 < len(arr) { // check if we've reached the end of the array
				if arr[i] > arr[i+1] {
					temp := arr[i]
					arr[i] = arr[i+1]
					arr[i+1] = temp
					sorted = false
				}
			}
		}
	}
	return arr
}

func main() {
	// silence is gold
	given := []int{1, 5, 4, 2, 3, 8, 6, 9, 7, 8}
	arr := bubbleSort(given)
	fmt.Println(arr)
}
