package main

import "fmt"

func main() {
	alphabets := [4]string{"1", "2", "3", "4"}
	x := alphabets[:]
	y := alphabets[:2]
	z := alphabets[1:]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}