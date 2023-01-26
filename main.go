package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world! \n")
	}
	http.HandleFunc("/hello", handler)
	log.Println("Listening to the localhost/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
