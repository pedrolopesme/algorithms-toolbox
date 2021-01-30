package main

func findBiggestNumber(arr []int) (biggest int) {
	for _, num := range arr {
		if num > biggest {
			biggest = num
		}
	}
	return
}

func RadixSort(arr []int) []int {
	biggest := findBiggestNumber(arr)
	size := len(arr)
	tempArr := make([]int, size, size)
	digit := 1

	for biggest/digit > 0 {
		buckets := [10]int{0}

		for i := 0; i < size; i++ {
			buckets[(arr[i]/digit)%10]++
		}

		for i := 1; i < 10; i++ {
			buckets[i] += buckets[i-1]
		}

		for i := size - 1; i >= 0; i-- {
			buckets[(arr[i]/digit)%10]--
			tempArr[buckets[(arr[i]/digit)%10]] = arr[i]
		}

		arr = tempArr
		digit *= 10
	}

	return arr
}
