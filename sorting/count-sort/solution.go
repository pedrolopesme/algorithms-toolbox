package main

func CountSortRunes(arr []rune) []rune {
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

func findMax(arr []int) int {
	var temp int

	temp = arr[0]

	for _, e := range arr {
		if temp < e {
			temp = e
		}
	}

	return temp
}

func CountSortInt(arr []int) []int {
	count := make([]int, findMax(arr))

	for i := 0; i < len(arr); i++ {
		count[arr[i]-1]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	outArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		position := count[arr[i]-1]
		outArr[position-1] = arr[i]
		count[arr[i]-1]--
	}

	return outArr
}
