package main

import "fmt"

func main() {
	var p *int
	i := 66
	fmt.Println("value", i)
	p = &i
	*p = 8
	fmt.Println("value", i)
}