package main

import "fmt"

type User struct {
	Name string

	age int
}

func main() {

	fmt.Println("working with structs")

	user := User{"himanshu", 45}

	fmt.Println("printing structs", user)

}
