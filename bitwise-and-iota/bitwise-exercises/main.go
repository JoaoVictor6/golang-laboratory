package main

import (
	"fmt"
)

var (
	byte1 uint = 60
	byte2 uint = 13
	somaAND uint
	somaOR uint
	somaXOR uint
	somaLeftShift uint
	somaRightShift uint
)

func main(){
	fmt.Printf("Values: %v | %v\n\n", byte1, byte2)

	somaAND = byte1 & byte2
	fmt.Println("AND(&) Operator")
	fmt.Printf("%d\n%b\n\n", somaAND, somaAND)

	somaOR = byte1 | byte2
	fmt.Println("OR(|) Operator")
	fmt.Printf("%d\n%b\n\n", somaOR, somaOR)

	somaXOR = byte1 ^ byte2
	fmt.Println("XOR(^) Operator")
	fmt.Printf("%d\n%b\n\n", somaOR, somaXOR)

	somaLeftShift = byte1 << byte2
	fmt.Println("Left Shift(<<) Operator")
	fmt.Printf("%d\n%b\n\n", somaLeftShift, somaLeftShift)

	somaRightShift = byte1 >> byte2
	fmt.Println("Left Shift(>>) Operator")
	fmt.Printf("%d\n%b\n\n", somaRightShift, somaRightShift)
}