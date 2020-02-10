package main

import "fmt"

func main() {
	fmt.Prnt("input : ")
	var name string
	var age int
	var height float32
	var weigth float32
	n, err := fmt.Scan(&name, &age, &weigth, &height)
}