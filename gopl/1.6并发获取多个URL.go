package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Println(time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	bytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	ch <- fmt.Sprint(bytes)
}
