package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{2, 4, 6}
	b := []int{2, 4, 6}
	fmt.Println(reflect.DeepEqual(a. b))
	c := []string{"hi", "hello"}
	d := []string{"helo", "hi"}
}