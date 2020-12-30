package main

import "fmt"

func main() {
	var myint int
	var myIntPointer *int

	myint = 12
	myIntPointer = &myint

	fmt.Println(*myIntPointer)
}
