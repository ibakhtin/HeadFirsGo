// +build OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
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
	channel <- len(body)
}

func main() {
	channel := make(chan int)
	urls := []string{
		"https://duckduckgo.com",
		"https://google.com",
		"https://ya.ru",
	}
	for _, url := range urls {
		go responseSize(url, channel)
	}
	for _ = range urls {
		fmt.Println(<-channel)
	}
}
