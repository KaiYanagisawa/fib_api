package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("fib/", handleControl)
	http.ListenAndServe(":8080", nil)
}