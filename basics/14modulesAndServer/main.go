package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("hello mod")

	router := mux.NewRouter()

	router.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))

}

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to home</h1>"))

}
