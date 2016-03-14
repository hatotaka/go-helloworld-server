package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = "1"

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!<br>\n")
	fmt.Fprintln(w, "version 1")
}

func handleEnv(w http.ResponseWriter, r *http.Request) {
	for _, v := range os.Environ() {
		fmt.Fprintln(w, v)
	}
}

func main() {
	http.HandleFunc("/", handleHelloWorld)
	http.HandleFunc("/env", handleEnv)

	http.ListenAndServe(":8080", nil)
}
