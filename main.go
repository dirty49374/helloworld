package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
		fmt.Println("received")
	})

	fmt.Println("listening on 80 ...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
