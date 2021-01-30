package main

func InsertionSort(arr []int) (output []int) {
	output = arr
	for i := range output {
		j := i
		for j > 0 && output[j-1] > output[j] {
			temp := output[j-1]
			output[j-1] = output[j]
			output[j] = temp
			j--
		}
	}
	return
}
