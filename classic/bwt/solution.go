package main

import "sort"

// generatePermutations computes the block
// permutations needed to the BWT
// TODO add tests
func generatePermutations(str string) (permutations []string) {
	i := 0
	for i < len(str) {
		permutations = append(permutations, str[i:]+str[:i])
		i++
	}
	return permutations
}

// Transform returns a Burrows-Wheeler transformation
func Transform(str string) (transform string) {
	permutations := generatePermutations(str)
	sort.Strings(permutations)

	for index := range permutations {
		permutation := permutations[index]
		transform += string(permutation[len(permutation)-1])
	}

	for index := range permutations {
		if permutations[index] == str {
			break
		}
	}

	return
}

// getIndexes returns the index of each mapped symbol
// from the sorted string
func getIndexes(str string, sorted string) (indexes []int) {
	usedPositions := make(map[int]bool)
	i := 0
	j := 0
	for i < len(str) {
		for j < len(sorted) {
			if sorted[i] == str[i] && !usedPositions[j] {
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
func InverseTransform() (tranform string) {
	return
}
