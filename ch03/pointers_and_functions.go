package main

import "fmt"

func createPointer() *float64 {
	var number = 36.6
	return &number
}

func printPointer(floatPointer *float64) {
	fmt.Println(*floatPointer)
}

func main() {
	printPointer(createPointer())
}