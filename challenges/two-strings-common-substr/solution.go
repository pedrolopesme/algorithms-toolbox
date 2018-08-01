package main

// Complete the twoStrings function below.
func createCharsArray(str string) map[byte]bool {
	chars := make(map[byte]bool, 26)
	for i := 0; i < len(str); i++ {
		chars[byte(str[i])] = true
	}
	return chars
}

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {
	var charsMap map[byte]bool
	var compareWith *string

	if len(s1) < len(s2) {
		charsMap = createCharsArray(s1)
		compareWith = &s2
	} else {
		charsMap = createCharsArray(s2)
		compareWith = &s1
	}
	for i := 0; i < len(*compareWith); i++ {
		if charsMap[byte((*compareWith)[i])] {
			return "YES"
		}
	}

	return "NO"
}
