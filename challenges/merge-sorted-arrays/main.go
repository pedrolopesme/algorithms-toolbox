package main

import "fmt"

func MergeSortedArrays(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}

	mergedArray := []int{}
	for len(arr1) > 0 || len(arr2) > 0 {
		if len(arr1) == 0 {
			mergedArray = append(mergedArray, arr2...)
			break
		} else if len(arr2) == 0 {
			mergedArray = append(mergedArray, arr1...)
			break
		} else if arr1[0] < arr2[0] {
			mergedArray = append(mergedArray, arr1[0])
			arr1 = arr1[1:]
		} else {
			mergedArray = append(mergedArray, arr2[0])
			arr2 = arr2[1:]
		}
	}

	return mergedArray
}

func main() {
	a1 := []int{1, 2}
	a2 := []int{3, 4}
	fmt.Println(MergeSortedArrays(a1, a2))
}
