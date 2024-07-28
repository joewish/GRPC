package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("server listening on 9000")
	}
}
