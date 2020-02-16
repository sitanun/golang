package main

import "fmt"

func main() {
	ch :=make(chan string, 2)
	ch <- "Hello"
	ch <- "Hi"
	fmt.Prinln(<-ch)
	fmt.Prinln(<-ch)
}