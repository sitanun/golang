package main

import "fmt"

type I interfeca{}

func main() {
	var i I
	i = "Hello"
	s, ok := i.(string)
}