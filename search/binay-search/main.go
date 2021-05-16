package main

func binarySearchRecursive(arr []int, val, from, to int) int {
	if from > to {
		return -1
	}

	mid := (from + to) / 2

	if arr[mid] == val {
		return mid
	} else if val < arr[mid] {
		return binarySearchRecursive(arr, val, from, mid-1)
	} else {
		return binarySearchRecursive(arr, val, mid+1, to)
	}
}
