package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	return merge(
		mergeSort(arr[:middle]),
		mergeSort(arr[middle:]))
}

func merge(left, right []int) []int {
	var newArr []int
	var leftIndex, rightIndex int

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			newArr = append(newArr, left[leftIndex])
			leftIndex++
		} else {
			newArr = append(newArr, right[rightIndex])
			rightIndex++
		}
	}
	newArr = append(newArr, left[leftIndex:]...)
	newArr = append(newArr, right[rightIndex:]...)
	return newArr
}

func main() {
}
