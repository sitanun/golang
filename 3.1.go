package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, I am %d years old\n", "sitanun", 19)

	fmt.Printf("type = %T \n", 2.89898989898)
	fmt.Printf("pi = %f \n", 2.89898989898)
	fmt.Printf("pi = %2.2f \n", 2.89898989898)
	fmt.Printf("pi = %9.f \n", 2.89898989898)
	fmt.Printf("pi = %-9.f \n", 2.89898989898)
	fmt.Printf("pi = %09.f \n", 2.89898989898)
	fmt.Printf("pi = %09.2f \n", 2.89898989898)
	fmt.Printf("pi = %E \n", 2.89898989898)

	fmt.Printf("2 > 1 = %t \n", 1 > 2)
	fmt.Printf("10 (base 2) = %b \n", 10)
	fmt.Printf("10 (base 8) = %o \n", 10)
	fmt.Printf("10 (base 10) = %d \n", 10)
	fmt.Printf("10 (base 16) = %x \n", 10)
}