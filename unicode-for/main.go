package main

import (
	"fmt"
)

func main(){
	init := 65
	for ; init <= 90; init++ {
		fmt.Printf("%#U\n", init)
	}
}