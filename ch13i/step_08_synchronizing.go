package main

import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	channelAbc := make(chan string)
	channelDef := make(chan string)
	go abc(channelAbc)
	go def(channelDef)
	fmt.Print(<-channelAbc)
	fmt.Print(<-channelDef)
	fmt.Print(<-channelAbc)
	fmt.Print(<-channelDef)
	fmt.Print(<-channelAbc)
	fmt.Print(<-channelDef)
	fmt.Println()
}
