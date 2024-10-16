package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the awesome site</h1>")
}

func main() {
	const port = ":5000"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlerFunc)

	fmt.Println("Running server on port", port)
	http.ListenAndServe(port, mux)
}
