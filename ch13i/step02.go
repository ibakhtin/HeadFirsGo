package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	responseSize("https://duckduckgo.com")
	responseSize("https://google.com")
	responseSize("https://ya.ru")
}
