package main

func StringReversal(str string) string {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}
	return newStr
}

func main() {
}
