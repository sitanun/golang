package main

import (
	"fmt"
)

func main(){
	a := []int{3, 5, 7}
	b := a
	fmt.Println(a, b)
	a[0] = 20
	fmt.Println(a, b)
}
