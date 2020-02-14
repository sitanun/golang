package main

import "fmt"

func main() {
	panic("Hello pacic")
	text := recover()
	fmt.Println(text)
}