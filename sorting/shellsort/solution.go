package main

func ShellSort(arr []int) []int {
	increment := len(arr) / 2
	for increment > 0 {
		for i := increment; i < len(arr); i++ {
			j := i
			temp := arr[i]

			for j >= increment && arr[j-increment] > temp {
				arr[j] = arr[j-increment]
				j -= increment
			}

			arr[j] = temp
		}
		if increment == 2 {
			increment = 1
		} else {
			increment = int(increment * 5 / 11)
		}
	}

	return arr
}