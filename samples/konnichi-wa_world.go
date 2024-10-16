// Bon: https://www.youtube.com/watch?v=_cmqniwQz3c
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
	err := http.ListenAndServe("localhost:12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
