package main

import "fmt"

func main() {
	alphabets := [2][3]string{{"a", "b", "c"}, {"A", "B", "C"}}
	fmt.Prinln(alphabets)
	fmt.Prinln(alphabets[0][1])
}