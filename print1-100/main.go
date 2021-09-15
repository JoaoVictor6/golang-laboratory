package main

import (
	"fmt"
)

func main(){
	for init := 1; init <= 10000; {
		fmt.Println(init)
		init += 1
	}
}