package main

type Fibo struct {
}

func (f *Fibo) CalculateIterative(index int) int {
	arrFibo := make([]int, index+1)
	for i := 0; i <= index; i++ {
		if i == 0 {
			arrFibo[0] = 0
		} else if i == 1 {
			arrFibo[1] = 1
		} else {
			arrFibo[i] = arrFibo[i-2] + arrFibo[i-1]
		}
	}
	return arrFibo[len(arrFibo)-1]
}

func main() {

}
