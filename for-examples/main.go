package main

import "fmt"

func main()  {
	var i, j, lenght int
	lenght = 10 
	for i=0; i<lenght;i++ {
		for j = 0; j < i; j++{
			fmt.Print("-")
		}
		fmt.Printf("\n")
	}

	for i=lenght; i>0;i-=1 {
		for j = i; j > 0; j-=1{
			fmt.Print("-")
		}
		fmt.Printf("\n")
	}
}