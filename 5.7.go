package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[1])
	numbers[1] = 8
	fmt.Println(numbers[1])
	length := len(numbers)
	fmt.Println("Lenght =", length)
}