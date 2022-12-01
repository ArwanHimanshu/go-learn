package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"` // replace key in json

	Price int

	Platform string `json:"website"`

	Password string `json:"-"` // remove the field from access

	Tags []string `json:"tags,omitempty"` // if value is nill then it remove the field from json
}

func main() {

	fmt.Println("working with json")
	DecodeJson()
}

func EncodingJson() {

	myCourses := []course{{"react bootcamp", 100, "web", "h123", []string{"react"}}, {"angular bootcamp", 100, "web", "h123", nil}}

	// finalJson, err := json.Marshal(myCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	
	{
		"coursename": "react bootcamp",
		"Price": 100,
		"website": "web",
		"tags": ["react"]
}

	`)

	var myCourses course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("json is valid")
		json.Unmarshal(jsonData, &myCourses)

		fmt.Printf("%#v", myCourses)

	} else {

		fmt.Println("json not valid")
	}

	// without struct only key value pair

	var myJson map[string]interface{}

	json.Unmarshal(jsonData, &myJson)

	fmt.Printf("%#v\n", myJson)

}
