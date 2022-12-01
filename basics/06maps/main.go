package main

import "fmt"

func main() {
	fmt.Println("working with time")

	myMaps := make(map[string]string)

	myMaps["JS"] = "Javascript"
	myMaps["PY"] = "python"

	fmt.Println("List of all ", myMaps)

	delete(myMaps, "JS")

	for key, value := range myMaps {

		fmt.Printf("key and value is %v and %v", key, value)

	}

	for _, value := range myMaps {

		fmt.Printf("value is %v", value)

	}
}
