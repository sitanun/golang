package main

import "fmt"

func main() {

	fmt.Printf( "20 is of type %T\n", 20)

	fmt.Printf( "20 is of type %T\n", string(20) )

	fmt.Printf( "20 is of type %T\n", float64(20) )
}