package main

import "fmt"

func main() {

	fmt.Println("working with defer")

	defer fmt.Println("one")

	defer fmt.Println("two")

	test()
}

func test() {

	for i := 0; i < 5; i++ {

		defer fmt.Println(i)
	}
}
