package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "app2\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8082", nil)
}
