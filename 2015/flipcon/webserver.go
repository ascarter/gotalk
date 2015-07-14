package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprintf(w, "Hello!\nURL = %s\n", req.URL)
}

func main() {
	fmt.Println("Listening on localhost:7777/hello")
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
