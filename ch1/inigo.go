package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Monoya")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}