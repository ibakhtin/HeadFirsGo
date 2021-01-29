package main

import "fmt"

func greeting(greetingChannel chan string) {
	greetingChannel <- "hi"
}

func main() {
	greetingChannel := make(chan string)
	go greeting(greetingChannel)
	receivedValue := <-greetingChannel
	fmt.Println(receivedValue)
}
