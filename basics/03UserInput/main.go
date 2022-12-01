package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input"

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating")

	// common, ok || commo, err;

	// input, _   ** ignore everything if anything goes wrong;
	// input, err   ** catch the error and handle it;
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating, ", input)
	fmt.Printf("type of the rating %T", input)

}
