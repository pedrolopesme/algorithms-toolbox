package main

import "math"

// HeapSort rearranges array using heap sort
func HeapSort(arr []int) ([]int) {
	arr = heapify(arr)
	end := len(arr) - 1

	for end > 0 {
		arr = swap(arr, end, 0)
		end--
		arr = siftDown(arr, 0, end)
	}
	return arr
}

// Swaps position i to j in arr array of integers
func swap(arr []int, i int, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr
}

// Heapify put elements of 'arr' in heap order, in-place
func heapify(arr []int) []int {
	start := iParent(len(arr))
	end := len(arr) - 1

	for start >= 0 {
		arr = siftDown(arr, start, end)
		start--
	}
	return arr
}

func siftDown(arr []int, start int, end int) [] int {
	root := start

	for iLeftChild(root) <= end {
		child := iLeftChild(root)
		swapElement := root

		if arr[swapElement] < arr[child] {
			swapElement = child
		}

		if child+1 <= end && arr[swapElement] < arr[child+1] {
			swapElement = child + 1
		}

		if swapElement == root {
			return arr
		} else {
			arr = swap(arr, root, swapElement)
			root = swapElement
		}
	}
	return arr
}

func iParent(i int) int {
	return int(math.Floor(float64(i-1) / 2))
}

func iLeftChild(i int) int {
	return 2*i + 1
}
