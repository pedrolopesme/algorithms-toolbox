package main

import (
	"sort"
)

// generatePermutations computes the block
// permutations needed to the BWT
func generatePermutations(str string) (permutations []string) {
	i := 0
	for i < len(str) {
		permutations = append(permutations, str[i:]+str[:i])
		i++
	}
	return permutations
}

// Transform returns a Burrows-Wheeler transformation
func Transform(str string) (index int, transform string) {
	permutations := generatePermutations(str)
	sort.Strings(permutations)

	for index = range permutations {
		permutation := permutations[index]
		transform += string(permutation[len(permutation)-1])
	}

	for index = range permutations {
		if permutations[index] == str {
			break
		}
	}

	return
}

// getIndexes returns the index of each mapped symbol
// from the sorted string
// todo add tests
func getIndexes(str string, sorted []string) (indexes []int) {
	usedPositions := make(map[int]bool)

	i := 0
	for i < len(str) {
		j := 0
		for j < len(sorted) {
			if sorted[j] == string(str[i]) && !usedPositions[j] {
				usedPositions[j] = true
				indexes = append(indexes, j)
			}
			j++
		}
		i++
	}
	return
}

// InverseTransform reverses a Burrows-Wheeler transform
// to the original string
// TODO implement and add tests
func InverseTransform(transformed string, originalIndex int) (original string) {
	var sorted []string
	i := 0
	for i < len(transformed) {
		sorted = append(sorted, string(transformed[i]))
		i++
	}
	sort.Strings(sorted)
	indexes := getIndexes(transformed, sorted)

	i = 0
	for i < len(transformed) {
		char := string(transformed[originalIndex])
		original = char + original
		originalIndex = indexes[originalIndex]
		i++
	}
	return
}
