package main

import (
	"fmt"
	"time"
)

func reportNam(name string, delay int) {
	for i := 1; i <= delay; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(name, "sleeping", i, "s")
	}
	fmt.Println(name, "wake up!")
}

func send(channel chan string) {
	reportNam("sending goroutine", 2)
	fmt.Println("<<< send() sending value A")
	channel <- "A"
	fmt.Println("<<< send() sending value B")
	channel <- "B"
}

func main() {
	channel := make(chan string)
	go send(channel)
	reportNam("receiving goroutine", 5)
	fmt.Println(">>> main() receiving value", <- channel)
	fmt.Println(">>> main() receiving value", <- channel)
}