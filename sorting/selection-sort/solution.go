package main

func SelectionSort(arr []int) ([]int) {
	for j := range arr[:len(arr) -1] {
		min := j
		for i := j + 1; i < len(arr); i++ {
			if arr[i] < arr[min] {
				min = i
			}
		}

		if min != j {
			temp := arr[j]
			arr[j] = arr[min]
			arr[min] = temp
		}
	}
	return arr
}