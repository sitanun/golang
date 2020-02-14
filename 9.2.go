package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var std [10]student
	std[0] = student{"sitanun", 19, "sitanun.ma"}

	fmt.Println(std[0])
}