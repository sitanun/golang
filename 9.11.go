package main

import "fmt"

type I interface{}

func desc(i I) {
	fmt.Printf("%v, %T \n", i, i)
}

func main() {
	var i I
	i = 3.1234586745
	desc(i)

}