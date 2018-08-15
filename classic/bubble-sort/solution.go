package main

func BubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr = swap(arr, i-1, i)
				swapped = true
			}
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
