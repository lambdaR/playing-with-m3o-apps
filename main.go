package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().String()
		greeting := "Apps branch Debo => " + t
		w.Write([]byte(greeting))
	})
	http.ListenAndServe(":8080", nil)
}
