package main

import "fmt"

type User struct {
	Name string

	age int
}

func (u User) GetStats() {
	fmt.Println("this is name", u.Name)

}

func main() {

	fmt.Println("working with structs")

	user := User{"himanshu", 45}

	user.GetStats()

	fmt.Println("printing structs", user)

	fmt.Println("printing structs", user)

}
