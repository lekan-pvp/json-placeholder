package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	r, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
