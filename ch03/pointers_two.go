package main

import "fmt"

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	*myIntPointer = 8
	fmt.Println(myInt)

	myFloat := 37.7
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	*myFloatPointer = 36.6
	fmt.Println(myFloat)
}
