package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E", "F"}
	deleteIndex := 5
	a = append(a[:deleteIndex])
}