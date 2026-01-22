package main

import (
	"fmt"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s", port)
	//fmt.Fprint(port, "Starting server on port %s")
	http.HandleFunc("/", HelloServer)
	_ = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path != "" {
		_, _ = fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	} else {
		_, _ = fmt.Fprint(w, "Hello World!")
	}
}
