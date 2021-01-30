package main

func CountSort(arr []rune) []rune {
	maxKey := 256
	count := make([]int, maxKey)
	for _, x := range arr {
		count[int(x)]++
	}

	accTotal := 0
	for i := 1; i < maxKey; i++ {
		oldCount := count[i]
		count[i] = accTotal
		accTotal += oldCount
	}

	output := make([]rune, maxKey)
	for _, x := range arr {
		output[count[int(x)]] = x
		count[int(x)]++
	}

	return output[:len(arr)]
}
