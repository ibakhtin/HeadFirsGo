// +build OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}

func main() {
	go responseSize("https://duckduckgo.com")
	go responseSize("https://google.com")
	go responseSize("https://ya.ru")
	time.Sleep(5 * time.Second)
}
