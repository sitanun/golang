package main

import "fmt"

func main() {
	fmt.Printlf("My name is %s, I am %d years old\n", "sitanun", 19)

	fmt.Printlf("type = %T \n", 2.8989898989)
	fmt.Printlf("pi = %f \n",2.8989898989)
	fmt.Printlf("pi = %2.2f \n",2.8989898989)
	fmt.Printlf("pi = %9.f \n", 2.8989898989)
	fmt.Printlf("pi = %-9.f \n", 2.8989898989)
	fmt.Printlf("pi = %09.f \n", 2.8989898989)
	fmt.Printlf("pi = %09.2f \n", 2.8989898989)
	fmt.Printlf("pi = %E \n", 2.8989898989)

	fmt.Printlf("2 > 1 =%t \n", 2 > 1)
	fmt.Printlf("20 (base 2) = %b \n", 20)
	fmt.Printlf("20 (base 4) = %o \n", 20)
	fmt.Printlf("20 (base 8) = %d \n", 20)
	fmt.Printlf("20 (base 10) = %x \n", 20)
}