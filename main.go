package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-Consumer-Username")
		fmt.Fprintf(w, "Hello World: %s", user)
	})
	http.HandleFunc("/teste", func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-Consumer-Username")
		fmt.Fprintf(w, "TESTEL: %s", user)
	})
	http.ListenAndServe(":8084", nil)
}
