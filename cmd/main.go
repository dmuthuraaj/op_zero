package main

import "net/http"

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	http.ListenAndServe(":8080", nil)
}
