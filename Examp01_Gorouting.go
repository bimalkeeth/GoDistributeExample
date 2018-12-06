package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server is starting on %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Hello world")
}
