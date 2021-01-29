package main

type FactorialRecursive struct {
}

func (f *FactorialRecursive) Calculate(number int) int {
	if number == 0 {
		return 1
	}
	return number * f.Calculate(number-1)
}

func main() {
	// silence is gold
}
