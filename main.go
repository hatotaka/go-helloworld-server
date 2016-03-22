package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var version = 14
var baseConfigmapVolume = "/etc/config"

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")
	fmt.Fprintf(w, "Version=%d\n", version)
	fmt.Fprintf(w, "env: HELLOWORLD_ENV=%s\n", os.Getenv("HELLOWORLD_ENV"))
	fmt.Fprintf(w, "configmap: helloworld.env=%s\n", getConfigMapVal("helloworld.env"))
}

func getConfigMapVal(path string) string {
	fullpath := filepath.Join(baseConfigmapVolume, path)

	if filepath.HasPrefix(fullpath, baseConfigmapVolume) != false {
		return ""
	}

	b, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return ""
	}

	return string(b)
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
