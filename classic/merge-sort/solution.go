package main

func MergeSort(a []int) []int {

	if len(a) <= 1 {
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)
	m := len(a) / 2

	for i, x := range a {
		switch {
		case i < m:
			left = append(left, x)
		case i >= m:
			right = append(right, x)
		}
	}

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {

	results := make([]int, 0)

	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				results = append(results, left[0])
				left = left[1:]
			} else {
				results = append(results, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			results = append(results, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			results = append(results, right[0])
			right = right[1:]
		}
	}

	return results
}
