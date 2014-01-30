package main

import (
	"net/http"
)

func main() {
	panic(http.ListenAndServe(":8081", http.FileServer(http.Dir("."))))
}
