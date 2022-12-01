package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rating")

	userInput, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", userInput)

	increaseRating, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)

	if err != nil {
		fmt.Println("error", err)

		panic(err)
	}

	fmt.Println("Thanks for rating, ", increaseRating+1)

}
