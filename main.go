package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = 4

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!<br>\n")
	fmt.Fprintf(w, "version %d<br>\n", version)

	env, notEmpty := os.LookupEnv("HELLOWORLD_ENV")
	if notEmpty {
		fmt.Fprintf(w, "HELLOWORLD_ENV=%d<br>\n", env)
	}else{
		fmt.Fprintf(w, "HELLOWORLD_ENV=null<br>\n", env)
	}

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
