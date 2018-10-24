package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", requestHandler)
	fmt.Println("Started listening...")
	http.ListenAndServe(":5000", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alive"))
}
