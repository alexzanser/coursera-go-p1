package main

import (
	"fmt"
	"log"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.URL.Path, "Launched test handler")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", testHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
