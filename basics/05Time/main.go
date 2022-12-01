package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Working with time")

	// date := time.Now()

	date := time.Date(2021, time.September, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println("current date is, ", date.Format("2006/02/01 Monday"))

}
