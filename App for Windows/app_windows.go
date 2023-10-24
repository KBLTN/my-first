package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi Kotya! version 2.0 \n"))

}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("Start HTTP server on port 14078")
	log.Fatal(http.ListenAndServe(":14078", nil))

}
