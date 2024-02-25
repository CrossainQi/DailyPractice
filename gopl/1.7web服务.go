package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count


func main() {
	http.HandleFunc("/", handler)
	http.HandlerFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprint(w, r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, count)
	mu.Unlock()
}
