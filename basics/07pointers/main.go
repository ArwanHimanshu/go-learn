package main

import "fmt"

func main() {

	fmt.Println("working with pointers")

	//declare a pointers
	// var ptr *int

	// fmt.Println("initial pointer value is", ptr)

	myNumber := 44
	var ptr = &myNumber

	fmt.Println("pointer's memory value is", ptr)

	//getting value
	fmt.Println("pointer's memory value is", *ptr)

	*ptr = *ptr * 2

	fmt.Println("pointer's memory value is", myNumber)

	var temp = add(myNumber)
	fmt.Println("pointer's memory value is", myNumber, temp)
}

func add(value int) int {
	value = value * 2
	return value
}
