package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}
