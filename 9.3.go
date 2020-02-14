package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var a student
	a.name = "Cartoon"
	a.age  = 19
	a.email = "sitanun.ma"

	b := student{"Cartoon", 4, "sitanun.ma"}

	c := student{name: "CC", email: "Abc@def"}

	d := student{age: 17}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}