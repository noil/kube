package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request for URL: " + r.URL.String())
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080", nil)
}
