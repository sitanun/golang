package main

import "fmt"

typt student struct {
	name  string
	age   int
	email string
}

func(std student) introduce() {
	fmt.Println("Hello my name is", std.name)
}