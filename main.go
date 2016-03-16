package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = 7

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!<br>\n")
	fmt.Fprintf(w, "Version=%d<br>\n", version)
	fmt.Fprintf(w, "HELLOWORLD_ENV=%s<br>\n", os.Getenv("HELLOWORLD_ENV"))

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
