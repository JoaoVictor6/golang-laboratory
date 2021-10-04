package main

func main() {
	SomaTudo([]int{}, []int{3, 9})
}

func Soma(arrayNumber []int) int {
	soma := 0
	// range iterator returns two values, array index and array value
	for _, numberValue := range arrayNumber {
		soma += numberValue
	}
	return soma
}

func SomaTudo(arrayNumbers ...[]int) (somas []int) {
	var soma []int
	for _, numbers := range arrayNumbers {
		soma = append(soma, Soma(numbers))
	}

	return soma
}
