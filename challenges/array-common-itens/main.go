package main

import "fmt"

func containsCommonItens(array1, array2 []rune) bool {
	array1Map := map[rune]bool{}
	for _, char := range array1 {
		array1Map[char] = true
	}

	for _, char := range array2 {
		if _, ok := array1Map[char]; !ok {
			return false
		}
	}
	return true
}

func main() {
	array1 := []rune{'a', 'b', 'c', 'd'}
	array2 := []rune{'z', 'y', 'a'}
	fmt.Println("Contains ", containsCommonItens(array1, array2))
}
