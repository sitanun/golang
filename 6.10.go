package main

import "fmt"

func main() {
	var a int
ReadInput:
	fmt.Print("type number :")
	fmt.Scan(&a)
	if a < 40 {
		goto ReadInput
	}
	fmt.Println(a)
}