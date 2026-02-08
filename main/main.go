package main

import (
	"fmt"
	"net/http"
	"redirect-server/redirect"
)

func main() {
	mux := defaultMux()

	pathsToURL := map[string]string{
		"/docs": "https://golang.org/doc/",
		"/yaml": "https://yaml.org/",
	}

	mapHandler := redirect.MapHandler(pathsToURL, mux)
	yaml := `
- path: /repo
  url: https://github.com/
- path: /golang
  url: https://golang.org
`

	yamlHandler, err := redirect.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
