package main

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	start, end := 0, len(arr)-1
	pivot := generatePivot(arr, start, end)
	arr[pivot], arr[end] = arr[end], arr[pivot]

	for i := range arr {
		if arr[i] < arr[end] {
			arr[start], arr[i] = arr[i], arr[start]
			start++
		}
	}

	arr[start], arr[end] = arr[end], arr[start]

	QuickSort(arr[:start])
	QuickSort(arr[start+1:])
	return arr
}

func generatePivot(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start
	for j := i; j < start-1; j++ {
		if arr[j] < pivot {
			arr = swap(arr, i, j)
			i++
		}
	}
	arr = swap(arr, i, end)
	return i
}

func swap(arr []int, i int, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr
}
