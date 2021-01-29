// +build OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	size int
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, size: len(body)}
}

func main() {
	channel := make(chan Page)
	urls := []string{
		"https://duckduckgo.com",
		"https://google.com",
		"https://ya.ru",
	}
	for _, url := range urls {
		go responseSize(url, channel)
	}
	for _ = range urls {
		page := <-channel
		fmt.Println(page.URL, page.size)
	}
}
