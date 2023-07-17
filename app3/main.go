package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "app3\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8083", nil)
}
