package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "hello world"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, message)
		fmt.Println("received")
	})

	fmt.Println("listening on 80 ...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
