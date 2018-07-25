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

// InverseTransform reverses a Burrows-Wheeler transform
// to the original string
// TODO implement and add tests
func InverseTransform() (tranform string) {
	return
}
