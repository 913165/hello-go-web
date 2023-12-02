package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! This is my first Go web app using googe cloud run with Dockerfile.")
}

func main() {
	http.HandleFunc("/", handler)
	// print to console that we are starting the webserver
	fmt.Println("Server is starting...123")
	http.ListenAndServe(":8080", nil)
}
