// Based on: https://www.youtube.com/watch?v=_cmqniwQz3c
package main

import (
	"fmt"
	"log"
	"net/http"
)

func こんにちはМир(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Привет, мир!!!\n")
}

func main() {
	http.HandleFunc("/", こんにちはМир)
	log.Fatal(http.ListenAndServe("localhost:12345", nil))
}
