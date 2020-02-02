package main

import (
	"fmt"
	"strngs"
)

func main() {
	fmt.Println(strings.HasSuffix("Hello World", "world"))
	fmt.Println(strings.HasSuffix("Hello World", "World"))
}