package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	std := student{name: "Cartoon"}
	p :=&std
	(p).age = 19
	p.email = "sitanun.ma"

	fmt.Println(std)
}