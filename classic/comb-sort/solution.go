package main

import "math"

func CombSort(arr []int) []int {
	gap := len(arr)
	shrinkRatio := 1.3
	sorted := false

	for !sorted {
		gap = int(math.Floor(float64(gap) / shrinkRatio))

		if gap > 1 {
			sorted = false
		} else {
			gap = 1
			sorted = true
		}

		i := 0
		for i+gap < len(arr) {
			if arr[i] > arr[i+gap] {
				arr = swap(arr, i, i+gap)
				sorted = false
			}
			i++
		}
	}

	return arr
}

func swap(arr []int, i int, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr
}
