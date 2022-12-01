package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("working with files")

	content := "this is testing file content"

	file, err := os.Create("./f.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(" file length is ", length)

	defer file.Close()

	readFile("./f.txt")
}

func readFile(fileName string) {

	data, err := ioutil.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("reding content ", string(data))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
