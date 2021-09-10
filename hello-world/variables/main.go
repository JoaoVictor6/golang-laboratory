package main

import "fmt"

func main(){
	//declare variables
	var str string
	var n, m int
	var mn float32

	//assign values
	str = "Hellow world"
	n = 10
	m = 50
	mn = 2.45

	fmt.Println("Value of str=", str)
	fmt.Println("Value of n=", n)
	fmt.Println("Value of m=", m)
	fmt.Println("VAlue of mn=", mn)

	//declare and assign values to variables
	var city string = "Londom"
	var x int = 100

	fmt.Println("Value of city=", city)
	fmt.Println("Value of x=", x)

	//declare variable with defining its type

	country := "DE"
	val := 15

	fmt.Println("Value of country=", country)
	fmt.Println("Value of val=", val)

	//declare multiple variables
	var (
		name string
		email string
		age int
	)

	name = "john"
	email = "john@email.com"
	age = 27

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}